package chart

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/DMXMax/mge/util"
)

const MAX_CHAOS = 9
const MIN_CHAOS = 0

// enum
type Odds int

// enum of Odds
const (
	Impossible Odds = iota
	NearlyImpossible
	VeryUnlikely
	Unlikely
	FiftyFifty
	Likely
	VeryLikely
	NearlyCertain
	Certain
)

type tFateChart map[Odds][9]int

// map of Odds to array of nine probabilities
var FateChart = tFateChart{
	Impossible:       {50, 25, 15, 10, 5, 5, 0, 0, -20},
	NearlyImpossible: {75, 50, 35, 25, 15, 10, 5, 5, 0},
	VeryUnlikely:     {85, 65, 50, 45, 25, 15, 10, 5, 5},
	Unlikely:         {90, 75, 55, 50, 35, 20, 15, 10, 5},
	FiftyFifty:       {95, 85, 75, 65, 50, 35, 25, 15, 10},
	Likely:           {100, 95, 90, 85, 75, 55, 50, 35, 25},
	VeryLikely:       {105, 95, 95, 90, 85, 75, 65, 50, 45},
	NearlyCertain:    {115, 100, 95, 95, 90, 80, 75, 55, 50},
	Certain:          {125, 110, 95, 95, 90, 85, 80, 65, 55},
}

type Result struct {
	RollOdds Odds
	Chaos    int
	Odds     int
	Roll     int
	Text     string
	Event    *util.Event
}

func (r *Result) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d: %s ", r.Roll, r.Text))

	if r.Event != nil {
		sb.WriteString(fmt.Sprintf("Event: %s", r.Event))
	}

	return sb.String()
}

func (f *tFateChart) RollOdds(o Odds, chaos int) *Result {
	chaos = max(min(chaos, MAX_CHAOS), MIN_CHAOS)

	odds := FateChart[o][MAX_CHAOS-chaos]
	roll := rand.Intn(100) + 1

	r := evaluate(odds, roll)
	r.RollOdds = o
	r.Odds = odds
	r.Chaos = chaos
	r.Roll = roll

	if roll%11 == 0 && roll/11 <= chaos {
		r.Event = util.GetEvent()
	}
	return r
}

func evaluate(odds, roll int) *Result {
	var r = new(Result)

	exy := odds / 5
	exn := ((100 - odds) / 5 * 4) + (odds + 1)

	switch {
	//top 20% is exceptional yes
	case roll <= exy:
		r.Text = "Exceptional Yes"
	case roll <= odds:
		r.Text = "Yes"
	//top 20% of the failure range
	case roll >= exn:
		r.Text = "Exceptional No"
	default:
		r.Text = "No"
	}

	return r

}
