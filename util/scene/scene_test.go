package scene

import (
	"testing"
)

func TestRollChaosDie_ExpectedScene(t *testing.T) {
	// Test multiple times to ensure we can get expected scenes
	// (roll > chaos + 1). With chaos 1, rolls 3-10 should be expected.
	expectedCount := 0
	trials := 100

	for i := 0; i < trials; i++ {
		result := RollChaosDie(1)
		if result == nil {
			t.Fatal("RollChaosDie returned nil")
		}
		if result.Roll < 1 || result.Roll > 10 {
			t.Errorf("roll out of range: %d", result.Roll)
		}
		if result.Roll > 2 {
			expectedCount++
			if result.SceneType != "expected" {
				t.Errorf("roll %d > (chaos 1 + 1) should be expected, got %s", result.Roll, result.SceneType)
			}
		}
	}

	// With chaos 1 and 100 trials, we should get some expected scenes
	if expectedCount == 0 {
		t.Error("expected at least one expected scene in 100 trials with chaos 1")
	}
}

func TestRollChaosDie_AlteredScene(t *testing.T) {
	// Test that odd rolls <= (chaos + 1) give altered scenes
	// We'll test with chaos 9 to maximize chances (all rolls 1-10 should trigger)
	alteredCount := 0
	trials := 100

	for i := 0; i < trials; i++ {
		result := RollChaosDie(9)
		if result == nil {
			t.Fatal("RollChaosDie returned nil")
		}
		if result.Roll < 1 || result.Roll > 10 {
			t.Errorf("roll out of range: %d", result.Roll)
		}

		// Check logic: roll <= (chaos + 1) AND odd -> altered
		if result.Roll <= 10 && result.Roll%2 == 1 {
			alteredCount++
			if result.SceneType != "altered" {
				t.Errorf("roll %d (odd, <= chaos 9 + 1) should be altered, got %s", result.Roll, result.SceneType)
			}
		}
	}

	// With chaos 9 and 100 trials, we should get some altered scenes
	if alteredCount == 0 {
		t.Error("expected at least one altered scene in 100 trials with chaos 9")
	}
}

func TestRollChaosDie_InterruptScene(t *testing.T) {
	// Test that even rolls <= (chaos + 1) give interrupt scenes
	// We'll test with chaos 9 to maximize chances (all rolls 1-10 should trigger)
	interruptCount := 0
	trials := 100

	for i := 0; i < trials; i++ {
		result := RollChaosDie(9)
		if result == nil {
			t.Fatal("RollChaosDie returned nil")
		}
		if result.Roll < 1 || result.Roll > 10 {
			t.Errorf("roll out of range: %d", result.Roll)
		}

		// Check logic: roll <= (chaos + 1) AND even -> interrupt
		if result.Roll <= 10 && result.Roll%2 == 0 {
			interruptCount++
			if result.SceneType != "interrupt" {
				t.Errorf("roll %d (even, <= chaos 9 + 1) should be interrupt, got %s", result.Roll, result.SceneType)
			}
		}
	}

	// With chaos 9 and 100 trials, we should get some interrupt scenes
	if interruptCount == 0 {
		t.Error("expected at least one interrupt scene in 100 trials with chaos 9")
	}
}

func TestRollChaosDie_ChaosBoundaries(t *testing.T) {
	tests := []struct {
		name  string
		chaos int
	}{
		{"chaos 0", 0},
		{"chaos 5", 5},
		{"chaos 10", 10},
		{"chaos negative", -1},
		{"chaos 20", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic regardless of chaos value
			result := RollChaosDie(tt.chaos)
			if result == nil {
				t.Fatal("RollChaosDie returned nil")
			}
			if result.Roll < 1 || result.Roll > 10 {
				t.Errorf("roll out of range: %d", result.Roll)
			}
			// Verify scene type is one of the valid values
			validTypes := map[string]bool{
				"expected":  true,
				"altered":   true,
				"interrupt": true,
			}
			if !validTypes[result.SceneType] {
				t.Errorf("invalid scene type: %s", result.SceneType)
			}
		})
	}
}

func TestRollChaosDie_DescriptionFormat(t *testing.T) {
	result := RollChaosDie(5)
	if result.Description == "" {
		t.Error("description should not be empty")
	}

	// Description should contain the roll and chaos values
	// We can't test exact format since roll is random, but we can verify it's not empty
	if len(result.Description) < 10 {
		t.Errorf("description seems too short: %q", result.Description)
	}
}

func TestGetSceneAdjustment_SingleAdjustment(t *testing.T) {
	// Run multiple times to ensure we get valid adjustments
	singleAdjustmentCount := 0
	trials := 100

	for i := 0; i < trials; i++ {
		adjustments := GetSceneAdjustment()
		if adjustments == nil {
			t.Fatal("GetSceneAdjustment returned nil")
		}
		if len(adjustments) == 0 {
			t.Error("GetSceneAdjustment returned empty slice")
		}
		if len(adjustments) > 2 {
			t.Errorf("GetSceneAdjustment returned too many adjustments: %d", len(adjustments))
		}

		// Verify adjustments are valid
		validAdjustments := map[string]bool{
			"Remove A Character":       true,
			"Add A Character":          true,
			"Reduce/Remove An Activity": true,
			"Increase An Activity":     true,
			"Remove An Object":         true,
			"Add An Object":            true,
		}

		for _, adj := range adjustments {
			if !validAdjustments[adj] {
				t.Errorf("invalid adjustment: %q", adj)
			}
		}

		if len(adjustments) == 1 {
			singleAdjustmentCount++
		}
	}

	// We should get some single adjustments in 100 trials
	if singleAdjustmentCount == 0 {
		t.Error("expected at least one single adjustment in 100 trials")
	}
}

func TestGetSceneAdjustment_TwoAdjustments(t *testing.T) {
	// Run multiple times to ensure we can get two adjustments
	twoAdjustmentCount := 0
	trials := 100

	for i := 0; i < trials; i++ {
		adjustments := GetSceneAdjustment()
		if adjustments == nil {
			t.Fatal("GetSceneAdjustment returned nil")
		}

		if len(adjustments) == 2 {
			twoAdjustmentCount++
			// Both should be valid
			validAdjustments := map[string]bool{
				"Remove A Character":       true,
				"Add A Character":          true,
				"Reduce/Remove An Activity": true,
				"Increase An Activity":     true,
				"Remove An Object":         true,
				"Add An Object":            true,
			}

			for _, adj := range adjustments {
				if !validAdjustments[adj] {
					t.Errorf("invalid adjustment in two-adjustment result: %q", adj)
				}
			}
		}
	}

	// We should get some two-adjustment results in 100 trials (7-10 on d10 = 40%)
	if twoAdjustmentCount == 0 {
		t.Error("expected at least one two-adjustment result in 100 trials")
	}
}

func TestGetSceneAdjustment_ValidAdjustments(t *testing.T) {
	// Verify all possible adjustments are valid strings
	validAdjustments := map[string]bool{
		"Remove A Character":       true,
		"Add A Character":          true,
		"Reduce/Remove An Activity": true,
		"Increase An Activity":     true,
		"Remove An Object":         true,
		"Add An Object":            true,
	}

	// Run many trials to try to get all possible adjustments
	seenAdjustments := make(map[string]bool)

	for i := 0; i < 500; i++ {
		adjustments := GetSceneAdjustment()
		for _, adj := range adjustments {
			if !validAdjustments[adj] {
				t.Errorf("invalid adjustment: %q", adj)
			}
			seenAdjustments[adj] = true
		}
	}

	// We should have seen at least some of the adjustments
	if len(seenAdjustments) == 0 {
		t.Error("did not see any adjustments in 500 trials")
	}
}

func TestRollResult_Structure(t *testing.T) {
	result := RollChaosDie(5)

	// Verify all fields are populated
	if result.Roll == 0 {
		t.Error("Roll should not be zero")
	}
	if result.SceneType == "" {
		t.Error("SceneType should not be empty")
	}
	if result.Description == "" {
		t.Error("Description should not be empty")
	}

	// Roll should be 1-10
	if result.Roll < 1 || result.Roll > 10 {
		t.Errorf("Roll should be 1-10, got %d", result.Roll)
	}
}

func TestSceneTypeDistribution(t *testing.T) {
	// Statistical test: with chaos 5, we should get a mix of scene types
	trials := 1000
	counts := map[string]int{
		"expected":  0,
		"altered":   0,
		"interrupt": 0,
	}

	for i := 0; i < trials; i++ {
		result := RollChaosDie(5)
		counts[result.SceneType]++
	}

	// With chaos 5 and d10:
	// - Rolls 1-5: 50% chance, split between altered (odd) and interrupt (even)
	// - Rolls 6-10: 50% expected
	// So we expect roughly: 50% expected, 25% altered, 25% interrupt

	// Allow for variance, just check we got all three types
	if counts["expected"] == 0 {
		t.Error("expected to see some expected scenes in 1000 trials")
	}
	if counts["altered"] == 0 {
		t.Error("expected to see some altered scenes in 1000 trials")
	}
	if counts["interrupt"] == 0 {
		t.Error("expected to see some interrupt scenes in 1000 trials")
	}

	// Total should equal trials
	total := counts["expected"] + counts["altered"] + counts["interrupt"]
	if total != trials {
		t.Errorf("total counts %d should equal trials %d", total, trials)
	}
}
