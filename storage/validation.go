// Package storage provides shared game data structures and operations
package storage

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// MinGameNameLength is the minimum allowed length for a game name
	MinGameNameLength = 3
	// MaxGameNameLength is the maximum allowed length for a game name
	MaxGameNameLength = 32
)

// ReservedNames is a list of names that cannot be used for games
var ReservedNames = []string{"current", "list"}

// ValidateGameName validates a game name according to the rules:
// - Minimum 3 characters
// - Maximum 32 characters
// - Only alphanumeric characters and single spaces
// - Cannot be a reserved name
func ValidateGameName(name string) error {
	// Trim whitespace
	name = strings.TrimSpace(name)

	// Normalize multiple consecutive spaces to single space
	name = regexp.MustCompile("  +").ReplaceAllString(name, " ")

	if name == "" {
		return fmt.Errorf("game name cannot be empty")
	}

	// Validate minimum length
	if len(name) < MinGameNameLength {
		return fmt.Errorf("game name must be at least %d characters", MinGameNameLength)
	}

	// Validate maximum length
	if len(name) > MaxGameNameLength {
		return fmt.Errorf("game name cannot be longer than %d characters", MaxGameNameLength)
	}

	// Validate characters (alphanumeric + spaces)
	matched, _ := regexp.MatchString("^[a-zA-Z0-9 ]+$", name)
	if !matched {
		return fmt.Errorf("game name can only contain letters, numbers, and spaces (a-z, A-Z, 0-9, space)")
	}

	// Block reserved names
	nameLower := strings.ToLower(name)
	for _, reserved := range ReservedNames {
		if nameLower == reserved {
			return fmt.Errorf("game cannot be named '%s'", name)
		}
	}

	return nil
}

// SanitizeGameName sanitizes a game name by trimming whitespace and normalizing spaces.
// This should be called before validation to ensure consistent input.
func SanitizeGameName(name string) string {
	// Trim whitespace
	name = strings.TrimSpace(name)

	// Normalize multiple consecutive spaces to single space
	name = regexp.MustCompile("  +").ReplaceAllString(name, " ")

	return name
}

