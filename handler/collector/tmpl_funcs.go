package collector

import (
	"time"
	"unicode"
	"unicode/utf8"
)

func formatDate(date time.Time) string {
	return date.Format("02/01/2006")
}

func formatHour(date time.Time) string {
	return date.Format("15:04")
}

func formatTime(date time.Time) string {
	return date.Format(time.RFC3339)
}

func truncateByWords(input string, maxWords int) string {
	processedWords := 0
	wordStarted := false
	for i := 0; i < len(input); {
		r, width := utf8.DecodeRuneInString(input[i:])
		if !unicode.IsSpace(r) {
			i += width
			wordStarted = true
			continue
		}

		if !wordStarted {
			i += width
			continue
		}

		wordStarted = false
		processedWords++
		if processedWords == maxWords {
			const ending = "..."
			if (i + len(ending)) >= len(input) {
				// Source string ending is shorter than "..."
				return input
			}

			return input[:i] + ending
		}

		i += width
	}

	return input
}
