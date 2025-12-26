// Package scene provides utilities for scene management in Mythic games.
package scene

import (
	"fmt"
	"math/rand"
)

// RollResult represents the result of rolling the Chaos Die for scene determination.
type RollResult struct {
	Roll        int    // The d10 roll result (1-10)
	SceneType   string // Scene type: "expected", "altered", or "interrupt"
	Description string // Human-readable description of the scene type
}

// RollChaosDie rolls a d10 and determines the scene type based on the chaos factor.
// According to Mythic GME rules:
// - If roll == chaos: Interrupted Scene
// - Else if roll <= chaos: Altered Scene
// - Otherwise: Expected Scene
func RollChaosDie(chaos int) *RollResult {
	// Roll 1d10
	roll := rand.Intn(10) + 1

	// Determine scene type
	var sceneType string
	var description string

	if roll == chaos {
		sceneType = "interrupt"
		description = fmt.Sprintf("Interrupted Scene (roll: %d, chaos: %d)", roll, chaos)
	} else if roll <= chaos {
		sceneType = "altered"
		description = fmt.Sprintf("Altered Scene (roll: %d, chaos: %d)", roll, chaos)
	} else {
		sceneType = "expected"
		description = fmt.Sprintf("Expected Scene (roll: %d, chaos: %d)", roll, chaos)
	}

	return &RollResult{
		Roll:        roll,
		SceneType:   sceneType,
		Description: description,
	}
}

