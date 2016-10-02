package schema

import "testing"

func TestNewPerson(t *testing.T) {
	person := NewPerson("Kévin Dunglas", "dunglas@gmail.com")

	if person.Type != "Person" {
		t.Error("The type has not been set.")
	}

	if person.Name != "Kévin Dunglas" {
		t.Error("The name has not been set.")
	}

	if person.Email != "dunglas@gmail.com" {
		t.Error("The email has not been set.")
	}
}
