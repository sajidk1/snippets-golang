package greetings

import (
    "testing"
	"regexp"
	"fmt"
)

// TestHelloName calls greetings.Hello with a name, checking 
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Sajid"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Sajid")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Sajid") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHellosNames(t *testing.T) {
	
    //TODO add test to check list of names appears in outputted map

    names := []string{"Sajid", "Bob", "Alice"}
	//want := regexp.MustCompile(`^(?!(Sajid|Bob|Alice)$)(\w*)`)
	msg, err := Hellos(names)
	
	fmt.Println(msg, err)
    //if !want.Match(msg) || err != nil {
    //    t.Fatalf(`Hello("Sajid") = %q, %v, want match for %#q, nil`, msg, err, want)
    //}
}

// TestHelloEmpty calls greetings.Hello with an empty string, 
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
