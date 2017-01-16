package schema

// Person stores the name and the email of a contributor.
type Person struct {
	JsonLd
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
}

// NewPerson initializes a new Person instance with some sensitive default values.
func NewPerson(name string, email string) *Person {
	return &Person{JsonLd: JsonLd{Type: "Person"}, Name: name, Email: email}
}
