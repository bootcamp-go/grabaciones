package main

import "fmt"

// NewStudent is a function that returns a new student
func NewStudent(name, entrydate string) Student {
	return Student{
		Name:      name,
		EntryDate: entrydate,
	}
}

// Student is a struct type that contains information about a student
type Student struct {
	// Name is a field that represents the name of the student
	Name      string
	// EntryDate is a field that represents the entry date of the student
	EntryDate string
}

// NewBootcamp is a function that returns a new bootcamp
func NewBootcamp(organization, language string) Bootcamp {
	return Bootcamp{
		Organization: organization,
		Language:     language,
	}
}

// Bootcamp is a struct type that contains information about a bootcamp
type Bootcamp struct {
	// Organization is a field that contains information about the bootcamp organization
	Organization string
	// Language is a field that contains information about the bootcamp programming language
	Language string
	// Students is a field that contains information about the bootcamp students
	Students []Student
}

func (b *Bootcamp) registerStudent(s Student) {
	(*b).Students = append((*b).Students, s)
}

func (b Bootcamp) printInfo() {
	fmt.Printf(
		"Welcome to the %s bootcamp, this bootcamp is language %s and we have %d students.\n",
		b.Organization,
		b.Language,
		len(b.Students),
	)
}

func main() {
	// Create a new bootcamp
	bc := NewBootcamp("MeLi", "Go")

	// Create some students
	s1 := NewStudent("John", "01/01/2021")
	s2 := NewStudent("Jane", "15/01/2021")
	s3 := NewStudent("Joe", "25/03/2021")

	// Register the students in the bootcamp
	bc.registerStudent(s1)
	bc.registerStudent(s2)
	bc.registerStudent(s3)

	// Show information about the bootcamp
	bc.printInfo()
}