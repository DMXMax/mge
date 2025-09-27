package plot

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/DMXMax/mge/util/theme"
)

type PlotPoint struct {
	Action      int    `json:"Action"`
	Tension     int    `json:"Tension"`
	Mystery     int    `json:"Mystery"`
	Social      int    `json:"Social"`
	Personal    int    `json:"Personal"`
	Description string `json:"Description"`
}

type PlotPointChart struct {
	PlotPoints []PlotPoint `json:"plot_points"`
}

// LoadChart reads the plot points JSON dataset and returns it as a Go struct.
func LoadChart() (*PlotPointChart, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to resolve caller information")
	}

	chartPath := filepath.Join(filepath.Dir(currentFile), "plot_points.json")

	data, err := os.ReadFile(chartPath)
	if err != nil {
		return nil, fmt.Errorf("read plot points: %w", err)
	}

	var rawEntries []struct {
		Text   string         `json:"text"`
		Ranges map[string]int `json:"ranges"`
	}

	if err := json.Unmarshal(data, &rawEntries); err != nil {
		var legacy struct {
			PlotPointChart PlotPointChart `json:"plot_point_chart"`
		}
		if legacyErr := json.Unmarshal(data, &legacy); legacyErr == nil && len(legacy.PlotPointChart.PlotPoints) > 0 {
			return &legacy.PlotPointChart, nil
		}
		return nil, fmt.Errorf("decode plot points: %w", err)
	}

	chart := &PlotPointChart{PlotPoints: make([]PlotPoint, 0, len(rawEntries))}
	for _, entry := range rawEntries {
		chart.PlotPoints = append(chart.PlotPoints, PlotPoint{
			Action:      entry.Ranges["Action"],
			Tension:     entry.Ranges["Tension"],
			Mystery:     entry.Ranges["Mystery"],
			Social:      entry.Ranges["Social"],
			Personal:    entry.Ranges["Personal"],
			Description: entry.Text,
		})
	}

	return chart, nil
}

// GetChartEntry returns the last plot point whose range for the provided theme
// is strictly less than the supplied roll value.
func (c *PlotPointChart) GetChartEntry(roll int, themeType theme.ThemeType) (*PlotPoint, error) {
	if c == nil {
		return nil, fmt.Errorf("plot point chart is nil")
	}
	if roll < 1 || roll > 100 {
		return nil, fmt.Errorf("roll out of range: %d", roll)
	}

	var lastMatch *PlotPoint
	for i := range c.PlotPoints {
		point := &c.PlotPoints[i]
		//log.Printf("Point: %v", point)
		value, err := rangeForTheme(point, themeType)
		if err != nil {
			return nil, err
		}
		if value == 0 {
			continue
		}
		if value >= roll {
			lastMatch = point
			break
		}
	}

	if lastMatch == nil {
		return nil, fmt.Errorf("no plot point found for roll %d and theme %s", roll, themeType)
	}

	return lastMatch, nil
}

func rangeForTheme(point *PlotPoint, themeType theme.ThemeType) (int, error) {
	switch themeType {
	case theme.ThemeAction:
		return point.Action, nil
	case theme.ThemeTension:
		return point.Tension, nil
	case theme.ThemeMystery:
		return point.Mystery, nil
	case theme.ThemeSocial:
		return point.Social, nil
	case theme.ThemePersonal:
		return point.Personal, nil
	default:
		return 0, fmt.Errorf("unknown theme %q", themeType)
	}
}
