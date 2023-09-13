package collector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		input          string
		limitWords     int
		expectedOutput string
	}{
		{"Im Jack from Vietnam", 3, "Im Jack from..."},
		{"Im Jack from Vietnam", 4, "Im Jack from Vietnam"},
		{"Im Jack from Vietnam", 5, "Im Jack from Vietnam"},
	}
	for _, test := range tests {
		output := truncateByWords(test.input, test.limitWords)
		assert.Equal(t, test.expectedOutput, output)
	}
}
