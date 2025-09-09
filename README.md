# mge — Mythic Game Emulator utilities in Go

A small Go module that implements a Fate Chart roller and event generator inspired by the Mythic Game Emulator. It provides a simple API to roll outcomes given odds and a chaos factor, and optionally attach randomly generated events.

## Features

- Fate Chart evaluation with nine odds levels and chaos factor clamping
- Result classification: Exceptional Yes / Yes / No / Exceptional No
- Event trigger on doubles constrained by chaos factor
- Event composition with focus, action, and subject

## Requirements

- Go 1.24+
- Module sets `go 1.24` and `toolchain go1.24.x` in `go.mod`.
  - If you prefer not to auto-download toolchains, run with `GOTOOLCHAIN=local`.

## Quick Start

Run the example app (defaults to 50/50 odds, chaos 6):

```bash
go run .
```

You should see output similar to:

```
likely - 42: Yes Event: NPC Action: Guide Power
```

Note: Output is random each run.

### CLI flags

- `-o` (odds): odds name or prefix (text only)
  - Names: `impossible`, `nearly impossible`, `very unlikely`, `unlikely`, `fifty fifty`, `likely`, `very likely`, `nearly certain`, `certain`
  - Prefixes: e.g., `-o unlikely`, `-o very`, `-o nearly` (ambiguous prefixes error)
- `-c` (chaos): integer chaos factor (0–8)

Examples:

```bash
go run . -o unlikely -c 5
go run . -o nearly certain -c 7
go run . -o very           # ambiguous: refine to 'very likely' or 'very unlikely'
```

## Project Structure

- `main.go`: Minimal CLI example that performs a single roll
- `chart/`: Fate Chart odds, evaluation logic, and tests
- `util/`: Event focus, action, and subject data and helpers

## Packages

### `chart`

- `type Odds`: Odds enum (`Impossible` … `Certain`).
- `var FateChart`: Map of `Odds` → `[9]int` thresholds by chaos index.
- `func (f *tFateChart) RollOdds(o Odds, chaos int) *Result`: Rolls 1–100, evaluates result, and attaches an event when appropriate.
- `func MatchOddsPrefix(prefix string) []Odds`: Returns odds whose names start with `prefix` (or all for `?` or no match).
- `type Result`: Structured result with `RollOdds`, `Chaos`, `Odds` (threshold), `Roll`, `Text`, and optional `Event`.

### `util`

- `func GetEvent() *Event`: Returns a random event composed of focus, action, and subject.
- `func GetEventFocus() EventFocus`: Randomly chooses the event focus with weighted ranges.
- `var Action []string`, `var Subject []string`: Word lists for event composition.

## Using the API

Example: roll at 50/50 odds with chaos 6

```go
res := chart.FateChart.RollOdds(chart.FiftyFifty, 6)
fmt.Println(res.String())
```

Match odds by prefix (case-sensitive):

```go
candidates := chart.MatchOddsPrefix("nearly") // NearlyImpossible, NearlyCertain
```

## Development

Run tests (uses a local build cache to avoid sandbox issues):

```bash
GOCACHE=$(pwd)/.gocache go test ./... -count=1
```

Vet:

```bash
GOCACHE=$(pwd)/.gocache go vet ./...
```

## Notes

- `RollOdds` clamps chaos to the supported range and triggers an event on numeric doubles (`11,22,…,99`) when `roll/11 <= chaos`.
- For reproducible tests, consider seeding `math/rand` with a fixed seed or adding a variant that accepts `*rand.Rand`.
