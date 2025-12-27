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
// - Roll > Chaos Factor: Expected Scene
// - Roll <= Chaos Factor AND roll is odd (1, 3, 5, 7, 9): Altered Scene
// - Roll <= Chaos Factor AND roll is even (2, 4, 6, 8): Interrupt Scene
func RollChaosDie(chaos int) *RollResult {
	// Roll 1d10
	roll := rand.Intn(10) + 1

	// Determine scene type
	var sceneType string
	var description string

	if roll > chaos {
		sceneType = "expected"
		description = fmt.Sprintf("Expected Scene (roll: %d, chaos: %d)", roll, chaos)
	} else if roll%2 == 1 { // Odd: 1, 3, 5, 7, 9
		sceneType = "altered"
		description = fmt.Sprintf("Altered Scene (roll: %d, chaos: %d)", roll, chaos)
	} else { // Even: 2, 4, 6, 8
		sceneType = "interrupt"
		description = fmt.Sprintf("Interrupted Scene (roll: %d, chaos: %d)", roll, chaos)
	}

	return &RollResult{
		Roll:        roll,
		SceneType:   sceneType,
		Description: description,
	}
}

// rollAdjustment rolls a single adjustment from the Scene Adjustment Table (1-6 only).
// This is a helper function used when rolling 7-10 (which requires 2 adjustments).
func rollAdjustment() string {
	roll := rand.Intn(6) + 1
	switch roll {
	case 1:
		return "Remove A Character"
	case 2:
		return "Add A Character"
	case 3:
		return "Reduce/Remove An Activity"
	case 4:
		return "Increase An Activity"
	case 5:
		return "Remove An Object"
	case 6:
		return "Add An Object"
	default:
		return "Unknown Adjustment"
	}
}

// GetSceneAdjustment rolls 1d10 on the Scene Adjustment Table and returns the result(s).
// According to Mythic GME rules:
// - 1: Remove A Character
// - 2: Add A Character
// - 3: Reduce/Remove An Activity
// - 4: Increase An Activity
// - 5: Remove An Object
// - 6: Add An Object
// - 7-10: Make 2 Adjustments (roll twice, ignoring results of 7-10)
// Returns a slice of strings, with one or two adjustment suggestions.
func GetSceneAdjustment() []string {
	roll := rand.Intn(10) + 1
	switch roll {
	case 1:
		return []string{"Remove A Character"}
	case 2:
		return []string{"Add A Character"}
	case 3:
		return []string{"Reduce/Remove An Activity"}
	case 4:
		return []string{"Increase An Activity"}
	case 5:
		return []string{"Remove An Object"}
	case 6:
		return []string{"Add An Object"}
	case 7, 8, 9, 10:
		// Make 2 adjustments - roll twice, ignoring 7-10
		adj1 := rollAdjustment()
		adj2 := rollAdjustment()
		return []string{adj1, adj2}
	default:
		return []string{"Unknown Adjustment"}
	}
}

