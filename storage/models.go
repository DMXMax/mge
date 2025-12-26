// Package storage provides shared game data structures and operations
// for Mythic applications, including Game, LogEntry, Thread, Character, and Scene models.
package storage

import (
	"time"

	"github.com/DMXMax/mge/util/theme"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LogEntry represents a single entry in a game's story log.
// Log entries can be of different types (e.g., dice rolls, story events)
// and are automatically timestamped.
type LogEntry struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      // When the log entry was created
	UpdatedAt time.Time      // When the log entry was last updated
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete support
	Type      int            // Type of log entry (0 = story, 1 = dice roll, etc.)
	Msg       string         // The log message content
	GameID    uuid.UUID      `gorm:"type:uuid"` // Foreign key to the game
}

// BeforeCreate is a GORM hook that generates a UUID for the log entry before creation.
func (l *LogEntry) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}

// Thread represents an adventure objective/goal in the Threads List.
// Threads can appear multiple times on the list (weighted, up to 3 times).
type Thread struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time      // When the thread was created
	UpdatedAt   time.Time      // When the thread was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index"`     // Soft delete support
	GameID      uuid.UUID      `gorm:"type:uuid"` // Foreign key to the game
	Name        string         // Name of the thread (required)
	Description string         // Optional description
	Weight      int            `gorm:"default:1"`      // How many times it appears on the list (1-3)
	Status      string         `gorm:"default:active"` // Status: "active", "resolved", "paused"
}

// BeforeCreate is a GORM hook that generates a UUID for the thread before creation.
func (t *Thread) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

// Character represents an important NPC in the Characters List.
// Characters can appear multiple times on the list (weighted, up to 3 times).
type Character struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time      // When the character was created
	UpdatedAt   time.Time      // When the character was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index"`     // Soft delete support
	GameID      uuid.UUID      `gorm:"type:uuid"` // Foreign key to the game
	Name        string         // Name of the character (required)
	Description string         // Optional description
	Weight      int            `gorm:"default:1"`      // How many times it appears on the list (1-3)
	Status      string         `gorm:"default:active"` // Status: "active", "inactive"
	Notes       string         // Additional notes about the character
}

// BeforeCreate is a GORM hook that generates a UUID for the character before creation.
func (c *Character) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

// Scene represents an active scene in a game.
// Scenes track the current narrative moment with its type (expected, altered, interrupt)
// and the expected concept for that scene.
type Scene struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt      time.Time      // When the scene was created
	UpdatedAt      time.Time      // When the scene was last updated
	DeletedAt      gorm.DeletedAt `gorm:"index"`     // Soft delete support
	GameID         uuid.UUID      `gorm:"type:uuid"` // Foreign key to the game
	Type           string         // Scene type: "expected", "altered", "interrupt"
	ExpectedConcept string        // The expected scene concept
	ChaosDieRoll   int           // The chaos die roll result
	IsActive       bool          `gorm:"default:true"` // Whether this scene is currently active
}

// BeforeCreate is a GORM hook that generates a UUID for the scene before creation.
func (s *Scene) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

// Game represents a Mythic game session with all its associated data.
// Each game has a name, chaos factor, story themes, and a log of events.
type Game struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt   time.Time      // When the game was created
	UpdatedAt   time.Time      // When the game was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index"`       // Soft delete support
	Name        string         `gorm:"uniqueIndex"` // Name of the game (unique)
	Chaos       int8           // Current Chaos level (1-9)
	StoryThemes theme.Themes   `gorm:"type:text"`         // Story themes for plot generation
	Log         []LogEntry     `gorm:"foreignKey:GameID"` // Associated log entries
	Threads     []Thread       `gorm:"foreignKey:GameID"` // Threads List
	Characters  []Character    `gorm:"foreignKey:GameID"` // Characters List
}

// BeforeCreate is a GORM hook that generates a UUID for the game before creation.
func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}

// SetChaos sets the chaos factor for the game.
// The chaos factor affects the likelihood of extreme results in dice rolls.
// Valid range is 1-9.
func (g *Game) SetChaos(v int8) {
	g.Chaos = v
}

// GetGameLog loads the most recent n log entries from the database into the game's Log field.
// The entries are ordered by creation date (newest first) and limited to n entries.
//
// Parameters:
//   - db: The database connection to use
//   - n: The number of log entries to load
//
// Returns an error if the database query fails.
func (g *Game) GetGameLog(db *gorm.DB, n int) error {
	// Query log entries directly from the database
	var entries []LogEntry
	result := db.Where("game_id = ?", g.ID).
		Order("created_at DESC").
		Limit(n).
		Find(&entries)

	if result.Error != nil {
		return result.Error
	}

	// Populate the game's Log field with the queried entries
	g.Log = entries

	return nil
}

