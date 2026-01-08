# mge - Mythic Game Emulator (Go Library)

## Project Overview

A Go library implementing the core Fate Chart mechanics and event generation for the **Mythic Game Master Emulator 2nd Edition**. This is the foundational library used by both `mythic-cli` and `mythic-mcp`.

## Purpose

Provides a clean, testable API for:
- Rolling on the Mythic Fate Chart with customizable odds and chaos
- Generating random events with focus, action, and subject
- Evaluating outcomes (Exceptional Yes, Yes, No, Exceptional No)
- Triggering events on dice doubles based on chaos factor

## Project Structure

```
mge/
├── chart/              # Fate Chart implementation
│   ├── chart.go       # Core chart logic and evaluation
│   ├── chaos.go       # Chaos factor management
│   ├── odds.go        # Odds enum and matching
│   └── *_test.go      # Test files
├── util/              # Event generation utilities
│   ├── event.go       # Event generation logic
│   ├── focus.go       # Event focus selection
│   └── data.go        # Action and Subject word lists
├── storage/           # Data structures for persistence (used by dependent projects)
├── main.go            # Example CLI application
├── go.mod             # Module definition
└── README.md          # User documentation
```

## Key Components

### Fate Chart (`chart/`)

**Odds Enum** (`chart/odds.go`):
```go
type Odds int

const (
    Impossible       Odds = 0
    NearlyImpossible Odds = 1
    VeryUnlikely     Odds = 2
    Unlikely         Odds = 3
    FiftyFifty       Odds = 4  // Default
    Likely           Odds = 5
    VeryLikely       Odds = 6
    NearlyCertain    Odds = 7
    Certain          Odds = 8
)
```

**Core Rolling Function**:
```go
// RollOdds rolls on the Fate Chart with given odds and chaos
// chaos is internal representation (0-8)
result := chart.FateChart.RollOdds(odds, chaos)
```

**Chaos Factor** (`chart/chaos.go`):
- User-facing range: 1-9
- Internal range: 0-8
- Convert using: `ChaosUserToInternal(userChaos)` and `ChaosInternalToUser(internalChaos)`

**Result Structure**:
```go
type Result struct {
    RollOdds int         // The odds threshold used
    Chaos    int         // Chaos factor (internal 0-8)
    Odds     int         // The threshold value
    Roll     int         // Actual d100 roll (1-100)
    Text     string      // "Yes", "No", "Exceptional Yes", etc.
    Event    *util.Event // Optional event on doubles
}
```

### Event Generation (`util/`)

**Event Structure**:
```go
type Event struct {
    Focus   EventFocus  // What the event is about
    Action  string      // Random action word
    Subject string      // Random subject word
    Meaning Meaning     // Descriptor and action pairs
}
```

**Event Focus** (`util/focus.go`):
- Weighted random selection (1-100)
- Focus types: Remote Event, NPC Action, Introduce NPC, Move Toward Thread, etc.
- Custom ranges per focus type

**Word Lists** (`util/data.go`):
- `Action []string` - Action words for event composition
- `Subject []string` - Subject words for event composition

### Storage Models (`storage/`)

Database models for games, logs, threads, characters (used by mythic-cli and mythic-mcp):
- `Game` - Game session data
- `LogEntry` - Story and dice roll logs
- `Thread` - Adventure objectives with weight
- `Character` - NPCs with weight

## Usage Examples

### Basic Fate Chart Roll

```go
package main

import (
    "fmt"
    "github.com/DMXMax/mge/chart"
)

func main() {
    // Roll at 50/50 odds with chaos 5
    chaos := chart.ChaosUserToInternal(5)
    result := chart.FateChart.RollOdds(chart.FiftyFifty, chaos)

    fmt.Println(result.String())
    // Output: "likely - 42: Yes" or similar
}
```

### Matching Odds by Prefix

```go
// Match odds by prefix (case-sensitive)
matches := chart.MatchOddsPrefix("nearly")
// Returns: [NearlyImpossible, NearlyCertain]

// Unambiguous match
matches := chart.MatchOddsPrefix("likely")
// Returns: [Likely]
```

### Generate Random Event

```go
package main

import (
    "fmt"
    "github.com/DMXMax/mge/util"
)

func main() {
    event := util.GetEvent()
    fmt.Printf("Focus: %s\n", util.EventText[event.Focus])
    fmt.Printf("Action: %s\n", event.Action)
    fmt.Printf("Subject: %s\n", event.Subject)
}
```

## Development

### Requirements

- Go 1.24+
- Module sets `go 1.24` and `toolchain go1.24.x` in `go.mod`
- If you prefer not to auto-download toolchains, run with `GOTOOLCHAIN=local`

### Building

```bash
# Build example CLI
go build -o mge

# Run example with custom odds and chaos
./mge -o likely -c 5
```

### Testing

```bash
# Run tests (uses local build cache)
GOCACHE=$(pwd)/.gocache go test ./... -count=1

# Vet code
GOCACHE=$(pwd)/.gocache go vet ./...

# Format code
go fmt ./...
```

### CLI Flags (Example App)

- `-o` (odds): Odds name or prefix (text only)
  - Names: `impossible`, `nearly impossible`, `very unlikely`, `unlikely`, `fifty fifty`, `likely`, `very likely`, `nearly certain`, `certain`
  - Prefixes: e.g., `-o unlikely`, `-o very`, `-o nearly` (ambiguous prefixes error)
- `-c` (chaos): Integer chaos factor (1-9)

Examples:
```bash
./mge -o unlikely -c 5
./mge -o "nearly certain" -c 7
./mge -o very  # Error: ambiguous (refine to 'very likely' or 'very unlikely')
```

## Important Implementation Details

### Event Triggering

Events trigger on numeric doubles (11, 22, 33, ..., 99) when:
```
roll / 11 <= chaos
```

Where chaos is the internal representation (0-8).

### Chaos Factor Clamping

The `RollOdds` function automatically clamps chaos to the valid internal range (0-8).

### Result String Format

The `Result.String()` method formats as:
```
{odds} - {roll}: {text} [Event: {focus} {action} {subject}]
```

Example: `"likely - 42: Yes Event: NPC Action: Guide Power"`

## Dependencies

- **Standard Library Only** - No external dependencies for core functionality
- `math/rand` - Random number generation
- Note: For reproducible tests, consider seeding `math/rand` or adding variants that accept `*rand.Rand`

## Integration with Other Projects

### mythic-cli

The CLI uses this library for:
- Fate Chart rolls in interactive shell
- Event generation for plot points
- Game state persistence (using storage models)

### mythic-mcp

The MCP server uses this library for:
- MCP tool implementations (`roll_fate_chart`, `generate_event`)
- Automatic logging to game sessions
- Thread/Character list integration with events

## Code Style Guidelines

- Follow standard Go conventions (gofmt, go vet)
- Keep packages focused (chart = fate chart, util = helpers)
- Use meaningful names (avoid abbreviations unless common)
- Document exported functions and types
- Write table-driven tests where appropriate

## Legal & Licensing

**Important:** This library implements copyrighted game mechanics from **Mythic Game Master Emulator 2nd Edition** by Tana Pigeon and Word Mill Games (© 2006, 2021, 2023).

- Users must own a legitimate copy of the official product
- Fan-created implementation for personal use
- Contact Word Mill Games for licensing questions: www.wordmillgames.com

## Quick Reference

### Common Types

```go
chart.Odds              // Enum for odds levels (0-8)
chart.Result            // Fate Chart roll result
util.Event              // Random event structure
util.EventFocus         // Event focus type
```

### Common Functions

```go
chart.FateChart.RollOdds(odds, chaos) *Result
chart.MatchOddsPrefix(prefix) []Odds
chart.ChaosUserToInternal(userChaos) int
chart.ChaosInternalToUser(internalChaos) int
util.GetEvent() *Event
util.GetEventFocus() EventFocus
```

### Constants

```go
chart.MinChaosUser = 1
chart.MaxChaosUser = 9
chart.MinChaosInternal = 0
chart.MaxChaosInternal = 8
```

## Testing Strategy

- Unit tests for odds matching
- Table-driven tests for chaos conversion
- Integration tests for full roll workflows
- Use local `.gocache` to avoid sandbox issues

## Notes for AI Assistants

When working on this library:
1. **Maintain backward compatibility** - This library is used by mythic-cli and mythic-mcp
2. **Test chaos conversion** - Internal (0-8) vs user-facing (1-9) can be confusing
3. **Event triggering logic** - Must follow official Mythic GME 2nd Edition rules
4. **No breaking changes** - Coordinate with dependent projects
5. **Keep dependencies minimal** - Standard library only preferred
6. **Document exported APIs** - This is a library, clear docs are critical
