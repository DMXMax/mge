package storage

import (
	"testing"
)

func TestValidateGameName_ValidNames(t *testing.T) {
	validNames := []string{
		"abc",
		"Test Game",
		"My Adventure 123",
		"The Quick Brown Fox",
		"a b c",
		"Game1",
		"ALLCAPS",
		"MixedCase123",
		"Exactly 32 Characters Long OK",
	}

	for _, name := range validNames {
		t.Run(name, func(t *testing.T) {
			err := ValidateGameName(name)
			if err != nil {
				t.Errorf("ValidateGameName(%q) returned error: %v", name, err)
			}
		})
	}
}

func TestValidateGameName_TooShort(t *testing.T) {
	shortNames := []string{
		"",
		"a",
		"ab",
		"  ",
		" a ",
	}

	for _, name := range shortNames {
		t.Run(name, func(t *testing.T) {
			err := ValidateGameName(name)
			if err == nil {
				t.Errorf("ValidateGameName(%q) should have returned error for too short", name)
			}
		})
	}
}

func TestValidateGameName_TooLong(t *testing.T) {
	longName := "This name is definitely way too long to be valid and exceeds maximum"
	err := ValidateGameName(longName)
	if err == nil {
		t.Errorf("ValidateGameName(%q) should have returned error for too long", longName)
	}
}

func TestValidateGameName_InvalidCharacters(t *testing.T) {
	invalidNames := []string{
		"Game!",
		"Test@Game",
		"My#Adventure",
		"Game$Name",
		"Test%",
		"Game^Name",
		"Test&Game",
		"Game*Name",
		"Test(Game)",
		"Game-Name",
		"Test_Game",
		"Game+Name",
		"Test=Game",
		"Game[Name]",
		"Test{Game}",
		"Game|Name",
		"Test\\Game",
		"Game:Name",
		"Test;Game",
		"Game'Name",
		"Test\"Game",
		"Game<Name>",
		"Test,Game",
		"Game.Name",
		"Test?Game",
		"Game/Name",
	}

	for _, name := range invalidNames {
		t.Run(name, func(t *testing.T) {
			err := ValidateGameName(name)
			if err == nil {
				t.Errorf("ValidateGameName(%q) should have returned error for invalid characters", name)
			}
		})
	}
}

func TestValidateGameName_ReservedNames(t *testing.T) {
	reservedNames := []string{
		"current",
		"list",
		"CURRENT",
		"LIST",
		"Current",
		"List",
	}

	for _, name := range reservedNames {
		t.Run(name, func(t *testing.T) {
			err := ValidateGameName(name)
			if err == nil {
				t.Errorf("ValidateGameName(%q) should have returned error for reserved name", name)
			}
		})
	}
}

func TestValidateGameName_WhitespaceHandling(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		shouldErr bool
	}{
		{"leading spaces", "  Test Game", false},
		{"trailing spaces", "Test Game  ", false},
		{"multiple spaces", "Test    Game", false},
		{"only spaces", "   ", true},
		{"tabs", "Test\tGame", true},
		{"newlines", "Test\nGame", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGameName(tt.input)
			if tt.shouldErr && err == nil {
				t.Errorf("ValidateGameName(%q) should have returned error", tt.input)
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("ValidateGameName(%q) returned unexpected error: %v", tt.input, err)
			}
		})
	}
}

func TestSanitizeGameName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Test Game", "Test Game"},
		{"  Test Game  ", "Test Game"},
		{"Test    Game", "Test Game"},
		{"  Multiple    Spaces  ", "Multiple Spaces"},
		{"", ""},
		{"   ", ""},
		{"NoSpaces", "NoSpaces"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := SanitizeGameName(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeGameName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSanitizeGameName_Integration(t *testing.T) {
	// Test that sanitize + validate works as expected
	input := "  Test    Game  "
	sanitized := SanitizeGameName(input)
	err := ValidateGameName(sanitized)
	if err != nil {
		t.Errorf("SanitizeGameName(%q) followed by ValidateGameName should succeed, got error: %v", input, err)
	}
}

func TestValidateGameName_ExactLengthBoundaries(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		shouldErr bool
	}{
		{"exactly 3 chars", "abc", false},
		{"exactly 32 chars", "12345678901234567890123456789012", false},
		{"2 chars", "ab", true},
		{"33 chars", "123456789012345678901234567890123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGameName(tt.input)
			if tt.shouldErr && err == nil {
				t.Errorf("ValidateGameName(%q) should have returned error", tt.input)
			}
			if !tt.shouldErr && err != nil {
				t.Errorf("ValidateGameName(%q) returned unexpected error: %v", tt.input, err)
			}
		})
	}
}
