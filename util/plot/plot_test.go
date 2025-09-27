package plot

import (
	"strings"
	"testing"

	"github.com/DMXMax/mge/util/theme"
)

func TestLoadChart(t *testing.T) {
	chart, err := LoadChart()
	if err != nil {
		t.Fatalf("LoadChart returned error: %v", err)
	}

	if chart == nil {
		t.Fatalf("LoadChart returned nil chart")
	}

	t.Logf("Actual Len: %d", len(chart.PlotPoints))
	if len(chart.PlotPoints) <= 1 {
		t.Fatalf("expected more than 1 plot point, got %d", len(chart.PlotPoints))
	}

	first := chart.PlotPoints[0]
	if first.Action != 8 || first.Tension != 8 || first.Mystery != 8 || first.Social != 8 || first.Personal != 8 {
		t.Errorf("unexpected theme values for first plot point: %+v", first)
	}
	if !strings.HasPrefix(first.Description, "CONCLUSION: If this Turning Point is currently a Plotline") {
		t.Errorf("first description prefix mismatch: %q", first.Description)
	}

	second := chart.PlotPoints[1]
	if second.Action != 24 || second.Tension != 24 || second.Mystery != 24 || second.Social != 24 || second.Personal != 24 {
		t.Errorf("unexpected theme values for second plot point: %+v", second)
	}
	if !strings.HasPrefix(second.Description, "NONE: Leave this Plot Point blank") {
		t.Errorf("second description prefix mismatch: %q", second.Description)
	}
}

func TestGetChartEntry(t *testing.T) {
	chart, err := LoadChart()
	if err != nil {
		t.Fatalf("LoadChart returned error: %v", err)
	}
	t.Run("returns last entry with range less than roll", func(t *testing.T) {
		entry, err := chart.GetChartEntry(90, theme.ThemeAction)
		if err != nil {
			t.Fatalf("GetChartEntry returned error: %v", err)
		}
		if entry == nil {
			t.Fatalf("GetChartEntry returned nil entry")
		}
		if entry.Description != "second" {
			t.Fatalf("expected \"second\" entry, got #%v", entry)
		}
	})

	t.Run("skips zero ranges", func(t *testing.T) {
		entry, err := chart.GetChartEntry(17, theme.ThemeTension)
		if err != nil {
			t.Fatalf("GetChartEntry returned error: %v", err)
		}
		if entry.Description != "second" {
			t.Fatalf("expected \"second\" entry, got %v", entry.Tension)
		}
	})

	t.Run("errors when no range is less than roll", func(t *testing.T) {
		if _, err := chart.GetChartEntry(-5, theme.ThemeAction); err == nil {
			t.Fatalf("expected error for roll with no matching range")
		}
	})

	t.Run("errors on invalid roll", func(t *testing.T) {
		if _, err := chart.GetChartEntry(0, theme.ThemeAction); err == nil {
			t.Fatalf("expected error for invalid roll")
		}
	})
}

func TestChartRangesAreIncreasing(t *testing.T) {
	chart, err := LoadChart()
	if err != nil {
		t.Fatalf("LoadChart returned error: %v", err)
	}

	themes := []struct {
		themeType theme.ThemeType
		name      string
		getValue  func(p *PlotPoint) int
	}{
		{theme.ThemeAction, "Action", func(p *PlotPoint) int { return p.Action }},
		{theme.ThemeTension, "Tension", func(p *PlotPoint) int { return p.Tension }},
		{theme.ThemeMystery, "Mystery", func(p *PlotPoint) int { return p.Mystery }},
		{theme.ThemeSocial, "Social", func(p *PlotPoint) int { return p.Social }},
		{theme.ThemePersonal, "Personal", func(p *PlotPoint) int { return p.Personal }},
	}

	lastValues := make(map[theme.ThemeType]int)
	for _, th := range themes {
		lastValues[th.themeType] = 0
	}

	for i, point := range chart.PlotPoints {
		for _, th := range themes {
			currentValue := th.getValue(&point)

			if currentValue == 0 {
				continue // Zeros are allowed, we just don't update lastValue.
			}

			if lastValue := lastValues[th.themeType]; currentValue <= lastValue {
				t.Errorf("point %d (%q): theme %s value %d is not greater than previous non-zero value %d", i+1, point.Description, th.name, currentValue, lastValue)
			}
			lastValues[th.themeType] = currentValue
		}
	}
}
