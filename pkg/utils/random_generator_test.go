package utils

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {

	tests := []struct {
		composition string
		length      int
	}{
		{"abc", 5},
		{"12345", 3},
		{"!@#", 4},
		{"", 0},
		{"abc123!@#", 6},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := RandomString(test.composition, test.length)
			if len(result) != test.length {
				t.Errorf("期望长度为 %d, 但得到 %d", test.length, len(result))
			}
			for _, char := range result {
				if !strings.ContainsRune(test.composition, char) {
					t.Errorf("字符 '%c' 不在组合中", char)
				}
			}
		})
	}
}
