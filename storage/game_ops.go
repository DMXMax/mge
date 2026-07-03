package storage

import (
	"fmt"

	"github.com/DMXMax/mge/util/scene"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteGamePermanently hard-deletes a game and all associated records.
// Uses Unscoped() so rows are removed from SQLite, not soft-deleted.
func DeleteGamePermanently(db *gorm.DB, gameID uuid.UUID) error {
	return db.Transaction(func(tx *gorm.DB) error {
		tables := []interface{}{
			&LogEntry{},
			&Scene{},
			&Thread{},
			&Character{},
		}
		for _, model := range tables {
			if err := tx.Unscoped().Where("game_id = ?", gameID).Delete(model).Error; err != nil {
				return fmt.Errorf("delete %T for game %s: %w", model, gameID, err)
			}
		}
		if err := tx.Unscoped().Where("id = ?", gameID).Delete(&Game{}).Error; err != nil {
			return fmt.Errorf("delete game %s: %w", gameID, err)
		}
		return nil
	})
}

// StartScene deactivates any active scene for the game and creates a new active scene.
// Must be called inside a transaction if combined with other writes.
func StartScene(db *gorm.DB, gameID uuid.UUID, concept string, rollResult *scene.RollResult) (*Scene, error) {
	if rollResult == nil {
		return nil, fmt.Errorf("rollResult is required")
	}

	if err := db.Model(&Scene{}).
		Where("game_id = ? AND is_active = ?", gameID, true).
		Update("is_active", false).Error; err != nil {
		return nil, fmt.Errorf("deactivate existing scene: %w", err)
	}

	newScene := Scene{
		GameID:          gameID,
		Type:            rollResult.SceneType,
		ExpectedConcept: concept,
		ChaosDieRoll:    rollResult.Roll,
		IsActive:        true,
	}
	if err := db.Create(&newScene).Error; err != nil {
		return nil, fmt.Errorf("create scene: %w", err)
	}
	return &newScene, nil
}
