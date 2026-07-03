package storage

import (
	"fmt"
	"testing"

	"github.com/DMXMax/mge/util/scene"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := InitDatabase(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("InitDatabase: %v", err)
	}
	if err := db.AutoMigrate(&Game{}, &LogEntry{}, &Thread{}, &Character{}, &Scene{}); err != nil {
		t.Fatalf("AutoMigrate: %v", err)
	}
	return db
}

func TestDeleteGamePermanently(t *testing.T) {
	db := setupTestDB(t)

	game := Game{Name: "Delete Me", Chaos: 4}
	if err := db.Create(&game).Error; err != nil {
		t.Fatalf("create game: %v", err)
	}

	log := LogEntry{Type: 0, Msg: "test log", GameID: game.ID}
	thread := Thread{GameID: game.ID, Name: "Thread A", Weight: 1}
	character := Character{GameID: game.ID, Name: "NPC A", Weight: 1}
	sceneRec := Scene{GameID: game.ID, Type: "expected", ExpectedConcept: "test", ChaosDieRoll: 5, IsActive: true}

	for _, rec := range []interface{}{&log, &thread, &character, &sceneRec} {
		if err := db.Create(rec).Error; err != nil {
			t.Fatalf("create related record: %v", err)
		}
	}

	if err := DeleteGamePermanently(db, game.ID); err != nil {
		t.Fatalf("DeleteGamePermanently: %v", err)
	}

	var count int64
	for _, model := range []interface{}{&Game{}, &LogEntry{}, &Thread{}, &Character{}, &Scene{}} {
		db.Unscoped().Model(model).Where("game_id = ?", game.ID).Count(&count)
		if count != 0 {
			t.Errorf("%T: expected 0 rows after delete, got %d", model, count)
		}
	}
	db.Unscoped().Model(&Game{}).Where("id = ?", game.ID).Count(&count)
	if count != 0 {
		t.Errorf("Game: expected 0 rows after delete, got %d", count)
	}
}

func TestDeleteGamePermanently_NotFound(t *testing.T) {
	db := setupTestDB(t)
	// Deleting a non-existent game should succeed (0 rows affected).
	if err := DeleteGamePermanently(db, uuid.New()); err != nil {
		t.Fatalf("DeleteGamePermanently on missing game: %v", err)
	}
}

func TestStartScene(t *testing.T) {
	db := setupTestDB(t)

	game := Game{Name: "Scene Game", Chaos: 4}
	if err := db.Create(&game).Error; err != nil {
		t.Fatalf("create game: %v", err)
	}

	oldScene := Scene{
		GameID: game.ID, Type: "expected", ExpectedConcept: "old",
		ChaosDieRoll: 8, IsActive: true,
	}
	if err := db.Create(&oldScene).Error; err != nil {
		t.Fatalf("create old scene: %v", err)
	}

	rollResult := &scene.RollResult{
		Roll: 3, SceneType: "altered", Description: "Altered Scene",
	}

	var created *Scene
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		created, err = StartScene(tx, game.ID, "new concept", rollResult)
		return err
	})
	if err != nil {
		t.Fatalf("StartScene: %v", err)
	}

	if created.Type != "altered" || created.ExpectedConcept != "new concept" || !created.IsActive {
		t.Errorf("unexpected created scene: %+v", created)
	}

	var old Scene
	if err := db.First(&old, oldScene.ID).Error; err != nil {
		t.Fatalf("reload old scene: %v", err)
	}
	if old.IsActive {
		t.Error("old scene should be deactivated")
	}

	var activeCount int64
	db.Model(&Scene{}).Where("game_id = ? AND is_active = ?", game.ID, true).Count(&activeCount)
	if activeCount != 1 {
		t.Errorf("expected 1 active scene, got %d", activeCount)
	}
}

func TestStartScene_RollbackOnFailure(t *testing.T) {
	db := setupTestDB(t)

	game := Game{Name: "Rollback Game", Chaos: 4}
	if err := db.Create(&game).Error; err != nil {
		t.Fatalf("create game: %v", err)
	}

	oldScene := Scene{
		GameID: game.ID, Type: "expected", ExpectedConcept: "keep",
		ChaosDieRoll: 9, IsActive: true,
	}
	if err := db.Create(&oldScene).Error; err != nil {
		t.Fatalf("create old scene: %v", err)
	}

	rollResult := &scene.RollResult{Roll: 2, SceneType: "interrupt", Description: "interrupt"}

	err := db.Transaction(func(tx *gorm.DB) error {
		if _, err := StartScene(tx, game.ID, "new", rollResult); err != nil {
			return err
		}
		// Force rollback
		return fmt.Errorf("simulated failure")
	})
	if err == nil {
		t.Fatal("expected transaction error")
	}

	var old Scene
	if err := db.First(&old, oldScene.ID).Error; err != nil {
		t.Fatalf("reload old scene: %v", err)
	}
	if !old.IsActive {
		t.Error("old scene should remain active after rollback")
	}
}
