package theme

import "testing"

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
	themes := Themes{ThemeAction, ThemeTension, ThemeMystery, ThemeSocial, ThemePersonal}
	counts := make(map[ThemeType]int)
	const trials = 100000

	for i := 0; i < trials; i++ {
		counts[themes.GetTheme()]++
	}
	t.Logf("%v, %v, %v, %v, %v", counts[themes[0]], counts[themes[1]], counts[themes[2]], counts[themes[3]], counts[themes[4]])

	for _, theme := range themeOrder {
		if counts[theme] == 0 {
			t.Fatalf("theme %q never selected", theme)
		}
	}

	if counts[ThemeAction] <= counts[ThemeMystery] {
		t.Fatalf("expected first theme count > third: got %d vs %d", counts[ThemeAction], counts[ThemeMystery])
	}
	if counts[ThemeAction] <= counts[ThemeTension] {
		t.Fatalf("expected first theme count > second: got %d vs %d", counts[ThemeAction], counts[ThemeTension])
	}
	if counts[ThemeTension] <= counts[ThemeMystery] {
		t.Fatalf("expected second theme count > third: got %d vs %d", counts[ThemeTension], counts[ThemeMystery])
	}
	lastCombined := counts[ThemeSocial] + counts[ThemePersonal]
	if lastCombined == 0 {
		t.Fatalf("expected fourth or fifth theme to appear")
	}
	if diff := counts[ThemeSocial] - counts[ThemePersonal]; diff < -trials/10 || diff > trials/10 {
		t.Fatalf("expected last two themes to be roughly even, diff=%d", diff)
	}
}
