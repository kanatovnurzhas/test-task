package service

import (
	"testing"
	"unicode"
)

func TestGenerateSalt(t *testing.T) {
	// Call the function to get a salt value
	salt := GenerateSalt()

	// Check that the length of the salt value is 12
	if len(salt) != 12 {
		t.Errorf("Expected salt length of 12, but got %d", len(salt))
	}

	// Check that the salt value contains only alphanumeric characters
	for _, char := range salt {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			t.Errorf("Salt contains non-alphanumeric character %q", char)
		}
	}
}
