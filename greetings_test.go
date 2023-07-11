package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"

	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Juan")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{
		"Juan",
		"David",
		"Rodrigo",
	}

	wantNameJuan := regexp.MustCompile(`\b` + names[0] + `\b`)
	wantNameDavid := regexp.MustCompile(`\b` + names[1] + `\b`)
	wantNameRodrigo := regexp.MustCompile(`\b` + names[2] + `\b`)

	messages, err := Hellos(names)

	if !wantNameJuan.MatchString(messages[names[0]]) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, messages[names[0]], err, wantNameJuan)
	}

	if !wantNameDavid.MatchString(messages[names[1]]) || err != nil {
		t.Fatalf(`Hello("David") = %q, %v, quiere coincidencia para %#q, nil`, messages[names[1]], err, wantNameDavid)
	}

	if !wantNameRodrigo.MatchString(messages[names[2]]) || err != nil {
		t.Fatalf(`Hello("Rodrigo") = %q, %v, quiere coincidencia para %#q, nil`, messages[names[2]], err, wantNameRodrigo)
	}
}
