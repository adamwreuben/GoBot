package core

import (
	"strings"
)

// - `?` matches exactly one occurrence of any character.
// - `*` matches arbitrary many (including zero) occurrences of any character.

type GoBotMatch struct {
	pattern []state
}

type state struct {
	NextChar    *rune
	HasWildcard bool
}

func (w *GoBotMatch) String() string {
	var sb strings.Builder
	for _, p := range w.pattern {
		if p.NextChar == nil {
			break
		}
		sb.WriteString(string(*p.NextChar))
	}
	return sb.String()
}

func NewMatch(pattern string) *GoBotMatch {
	simplified := make([]state, 0, len(pattern))
	prevWasStar := false
	for _, currentChar := range pattern {
		copyCurrentChar := currentChar
		if currentChar == '*' {
			prevWasStar = true
		} else {
			s := state{
				NextChar:    &copyCurrentChar,
				HasWildcard: prevWasStar,
			}
			simplified = append(simplified, s)
			prevWasStar = false
		}
	}

	if len(pattern) > 0 {
		final := state{
			NextChar:    nil,
			HasWildcard: prevWasStar,
		}
		simplified = append(simplified, final)
	}

	return &GoBotMatch{
		pattern: simplified,
	}
}

// Matches indicates whether the matcher finds a match in the input string.
func (w *GoBotMatch) Matches(input string) bool {
	if len(w.pattern) == 0 {
		return false
	}

	patternIdx := 0
	for _, inputChar := range input {

		if patternIdx > len(w.pattern) {
			return false
		}

		p := w.pattern[patternIdx]

		if p.NextChar != nil && (*p.NextChar == '?' || *p.NextChar == inputChar) {
			patternIdx += 1
		} else if p.HasWildcard {
			if p.NextChar == nil {
				return true
			}
		} else {
			// Go back to last state with wildcard
			for {
				pattern := w.pattern[patternIdx]
				if pattern.HasWildcard {
					if pattern.NextChar != nil && (*pattern.NextChar == '?' || *pattern.NextChar == inputChar) {
						patternIdx += 1
					}
					break
				}
				if patternIdx == 0 {
					return false
				}
				patternIdx -= 1
			}
		}
	}
	return w.pattern[patternIdx].NextChar == nil
}
