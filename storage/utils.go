// Package storage provides shared game data structures and operations
package storage

import (
	"strings"
)

// SanitizeFilename sanitizes a string to be safe for use as a filename.
// It replaces path separators and problematic characters with hyphens
// and collapses whitespace sequences.
func SanitizeFilename(s string) string {
	s = strings.TrimSpace(s)
	// Replace path separators and problematic characters
	repl := []string{"/", "-", "\\", "-", ":", "-", "*", "-", "?", "-", "\"", "-", "<", "-", ">", "-", "|", "-"}
	r := strings.NewReplacer(repl...)
	s = r.Replace(s)
	// Collapse whitespace
	s = strings.Join(strings.Fields(s), " ")
	return s
}

