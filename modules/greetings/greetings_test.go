package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Sajid"
	want := regexp.MustCompile(`\b` + name + `\b`)
	greeting, err := Hello("Sajid")
	if !want.MatchString(greeting) || err != nil {
		t.Fatalf(`Hello("Sajid") = %q, %v, want match for %#q, nil`, greeting, err, want)
	}
}

func TestHellosNames(t *testing.T) {

	// Names to test function with
	names := []string{"Sajid", "Bob", "Alice"}

	greetings, err := Hellos(names)

	// For each name we input, we expect a greeting back, so let's ensure that works by comparing the length
	namesLength := len(names)
	greetingsLength := len(greetings)
	if namesLength != greetingsLength || err != nil {
		t.Fatalf(`names length = %v, greetings length = %v, greetings = %q, error = %v`, namesLength, greetingsLength, greetings, err)
	}

	// Loop over the inputted list of names, using each name as a key to acccess its corresponding value in the map.
	// The corresponding value (i.e. greeting) should include our inputted name
	for _, name := range names {

		want := regexp.MustCompile(`\b` + name + `\b`)

		if !want.MatchString(greetings[name]) || err != nil {
			t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, greetings[name], err, want)
		}

	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
