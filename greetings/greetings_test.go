package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName вызывает greetings.Hello с именем, проверяя
// на допустимость возвращаемого значения.
func TestHelloName(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Alex")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Alex") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty вызывает greetings.Hello с пустой строкой,
// проверяя на наличие ошибки.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
