package theme

import (
	"math/rand/v2"
	"time"
)

type ThemeType string
type Themes [5]ThemeType

const (
	ThemeAction   ThemeType = "Action"
	ThemeTension  ThemeType = "Tension"
	ThemeMystery  ThemeType = "Mystery"
	ThemeSocial   ThemeType = "Social"
	ThemePersonal ThemeType = "Personal"
)

func (t ThemeType) String() string {
	return string(t)
}

// this function returns an array of five theme types.
// The types are chosen randomly, with equal probability.
// If the theme has already been chosen, chose the next in the list.
// If the end of the list is reached, start from the beginning.
// The list is: ThemeAction, ThemeTension, ThemeMystery, ThemeSocial, ThemePersonal.
func GetThemes() Themes {
	themes := Themes{ThemeAction, ThemeTension, ThemeMystery, ThemeSocial, ThemePersonal}
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	r.Shuffle(len(themes), func(i, j int) { themes[i], themes[j] = themes[j], themes[i] })
	return themes
}

func (ts Themes) GetTheme() ThemeType {
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	roll := r.IntN(10) + 1

	switch {
	case roll <= 4:
		return ts[0]
	case roll <= 7:
		return ts[1]
	case roll <= 9:
		return ts[2]
	default:
		if r.IntN(2) == 0 {
			return ts[3]
		}
		return ts[4]
	}
}
