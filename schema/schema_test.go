package schema

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	person := NewPerson("Kévin Dunglas", "dunglas@gmail.com")

	if person.Type != "Person" {
		t.Error("The type has not been set.")
	}

	if person.Context != "" {
		t.Error("The context should not be set directly on Person.")
	}

	if person.Name != "Kévin Dunglas" {
		t.Error("The name has not been set.")
	}

	if person.Email != "dunglas@gmail.com" {
		t.Error("The email has not been set.")
	}
}

func TestNewCreativeWork(t *testing.T) {
	creativeWork := NewCreativeWork()

	if creativeWork.Type != "CreativeWork" {
		t.Error("The type must be \"CreativeWork\".")
	}

	if creativeWork.Context != "http://schema.org" {
		t.Error("The type must be Schema.org.")
	}
}
