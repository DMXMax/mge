package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "time"

    "math/rand"

    "github.com/DMXMax/mge/chart"
)

//create a function called main that generates a random result from the Actions map and a random result from the subject map and prints them out
//hint: use the rand package
//hint: use the len function
//hint: use the rand.Intn function
//hint: use the fmt.Println function
//hint: use the util.Action map
//hint: use the util.Subject map

func main() {
    // Flags: -o for odds (name/prefix or index 0-8), -c for chaos
    oddsFlag := flag.String("o", "fifty", "odds name or prefix (e.g., 'unlikely', 'very', 'nearly certain')")
    chaos := flag.Int("c", 6, "chaos factor (0-8)")
    flag.Parse()

    rand.Seed(time.Now().UnixNano())

    o, err := parseOdds(*oddsFlag)
    if err != nil {
        fmt.Fprintf(os.Stderr, "invalid -o value: %v\n", err)
        os.Exit(2)
    }

    if *chaos < 0 || *chaos > 8 {
        fmt.Fprintf(os.Stderr, "invalid -c value: %d (must be 0..8)\n", *chaos)
        os.Exit(2)
    }

    result := chart.FateChart.RollOdds(o, *chaos)
    fmt.Printf("%s\n", result)
}

func parseOdds(s string) (chart.Odds, error) {
    s = strings.TrimSpace(strings.ToLower(s))
    if s == "" {
        return chart.FiftyFifty, nil
    }
    all := []chart.Odds{
        chart.Impossible, chart.NearlyImpossible, chart.VeryUnlikely, chart.Unlikely,
        chart.FiftyFifty, chart.Likely, chart.VeryLikely, chart.NearlyCertain, chart.Certain,
    }
    matches := make([]chart.Odds, 0, len(all))
    for _, o := range all {
        if strings.HasPrefix(o.String(), s) {
            matches = append(matches, o)
        }
    }
    if len(matches) == 0 {
        return 0, fmt.Errorf("no odds matched %q", s)
    }
    if len(matches) > 1 {
        for _, m := range matches {
            if m.String() == s {
                return m, nil
            }
        }
        return 0, fmt.Errorf("ambiguous odds %q; matches %d values, use a longer prefix", s, len(matches))
    }
    return matches[0], nil
}

/*func properTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = cases.Title(word)
		}
	}
	return strings.Join(words, " ")
}*/
