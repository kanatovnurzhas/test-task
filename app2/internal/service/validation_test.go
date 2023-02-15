package service

import "testing"

var tests = []struct {
	email string
	exp   bool
}{
	{"nurzhas@gmail.com", true},
	{"dias@mail.ru", true},
	{"", false},
	{"нуржас", false},
	{"nurzhas", false},
	{"  ", false},
	{"....@gmail.com", false},
	{"Nurzhas@gmail.com", false},
	{"nurzhas@gmail_com", false},
}

func TestValidateEmail(t *testing.T) {
	for _, val := range tests {
		res := ValidateEmail(val.email)
		if res != val.exp {
			t.Errorf("ValidateEmail(%v) = %v, expected %v", val.email, res, val.exp)
		}
	}
}
