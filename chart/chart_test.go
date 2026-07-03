package chart

import "testing"

func TestTriggersEventOnDouble(t *testing.T) {
	tests := []struct {
		name          string
		roll          int
		internalChaos int
		wantEvent     bool
	}{
		// user chaos 1 (internal 0): only 11 triggers
		{name: "chaos1 roll11", roll: 11, internalChaos: 0, wantEvent: true},
		{name: "chaos1 roll22", roll: 22, internalChaos: 0, wantEvent: false},

		// user chaos 5 (internal 4): 11–55 trigger, 66 does not
		{name: "chaos5 roll11", roll: 11, internalChaos: 4, wantEvent: true},
		{name: "chaos5 roll55", roll: 55, internalChaos: 4, wantEvent: true},
		{name: "chaos5 roll66", roll: 66, internalChaos: 4, wantEvent: false},

		// user chaos 9 (internal 8): all doubles through 99 trigger
		{name: "chaos9 roll99", roll: 99, internalChaos: 8, wantEvent: true},
		{name: "chaos9 roll88", roll: 88, internalChaos: 8, wantEvent: true},

		// non-doubles never trigger
		{name: "non-double roll12", roll: 12, internalChaos: 8, wantEvent: false},
		{name: "non-double roll50", roll: 50, internalChaos: 8, wantEvent: false},

		// clamping: internal chaos below/above range
		{name: "clamp low chaos1 roll11", roll: 11, internalChaos: -5, wantEvent: true},
		{name: "clamp high chaos9 roll99", roll: 99, internalChaos: 99, wantEvent: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := triggersEventOnDouble(tt.roll, tt.internalChaos)
			if got != tt.wantEvent {
				t.Errorf("triggersEventOnDouble(%d, %d) = %v, want %v",
					tt.roll, tt.internalChaos, got, tt.wantEvent)
			}
		})
	}
}
