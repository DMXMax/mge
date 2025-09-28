package theme

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetThemes(t *testing.T) {
	themes := GetThemes()
	t.Logf("themes: %v", themes)

	if len(themes) != len(themeOrder) {
		t.Fatalf("expected %d themes, got %d", len(themeOrder), len(themes))
	}

	seen := make(map[ThemeType]int, len(themeOrder))
	for _, theme := range themes {
		seen[theme]++
		if seen[theme] > 1 {
			t.Fatalf("duplicate theme %q returned", theme)
		}
	}

	for _, want := range themeOrder {
		if seen[want] != 1 {
			t.Fatalf("expected theme %q to appear exactly once", want)
		}
	}
}

var themeOrder = []ThemeType{
	ThemeAction,
	ThemeTension,
	ThemeMystery,
	ThemeSocial,
	ThemePersonal,
}

func TestThemesGetTheme(t *testing.T) {
	themes := GetThemes()
	counts := make(map[ThemeType]int)
	const trials = 100000

	for i := 0; i < trials; i++ {
		counts[themes.GetRandomTheme()]++
	}

	// Use a strings.Builder for efficient string construction.
	var sb strings.Builder
	// Start with a newline for better layout in the test log.
	sb.WriteString("\n")
	// Loop through the themes to build the output string.
	for i, theme := range themes {
		// Format each theme and its corresponding count, padding for alignment.
		sb.WriteString(fmt.Sprintf("%-10s %-7d", theme, counts[theme]))
		// Add a newline after every second item to create two columns,
		// or if it's the last item. Otherwise, add tabs for spacing.
		if i%2 != 0 || i == len(themes)-1 {
			sb.WriteString("\n")
		} else {
			sb.WriteString("\t\t")
		}
	}
	// Log the final, formatted string. This is more readable and maintainable
	// than a single complex format string.
	t.Log(sb.String())

	for _, theme := range themeOrder {
		if counts[theme] == 0 {
			t.Fatalf("theme %q never selected", theme)
		}
	}

	if counts[themes[0]] <= counts[themes[2]] {
		t.Fatalf("expected first theme count > third: got %d vs %d", counts[themes[0]], counts[themes[2]])
	}
	if counts[themes[0]] <= counts[themes[1]] {
		t.Fatalf("expected first theme count > second: got %d vs %d", counts[themes[0]], counts[themes[1]])
	}
	if counts[themes[1]] <= counts[themes[2]] {
		t.Fatalf("expected second theme count > third: got %d vs %d", counts[themes[1]], counts[themes[2]])
	}
	lastCombined := counts[themes[3]] + counts[themes[4]]
	if lastCombined == 0 {
		t.Fatalf("expected fourth or fifth theme to appear")
	}
	if diff := counts[themes[3]] - counts[themes[4]]; diff < -trials/10 || diff > trials/10 {
		t.Fatalf("expected last two themes to be roughly even, diff=%d", diff)
	}
}
