package schema

type Person struct {
	JsonLd
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
}

func NewPerson(name string, email string) *Person {
	return &Person{JsonLd: JsonLd{Type: "Person"}, Name: name, Email: email}
}
