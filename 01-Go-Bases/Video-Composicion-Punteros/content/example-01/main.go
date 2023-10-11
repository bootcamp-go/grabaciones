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
func (p *Person) SetFirstName(firstName string) {
	(*p).FirstName = firstName
}

func main() {
	// Create a new person
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	// update the first name
	person.SetFirstName("Jane")

	// print the full name
	println(person.FullName())
}