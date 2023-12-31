package main

// Person is a struct that represents a person
type Person struct {
	// FirstName is the first name of the person
	FirstName string
	// LastName is the last name of the person
	LastName  string
}

// FullName returns the full name of the person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// SetFirstName sets the first name of the person
func (p Person) SetFirstName(name string) {
	p.FirstName = name
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	person.SetFirstName("Jane")

	println(person.FullName())
}