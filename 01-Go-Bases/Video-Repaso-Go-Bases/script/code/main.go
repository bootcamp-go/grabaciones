package main

import "fmt"

// NewStudent is a function that creates a new student
func NewStudent(name string, entryDate string) *Student {
	return &Student{
		Name:      name,
		EntryDate: entryDate,
	}
}

// Student is a struct type that contains information about a student
type Student struct {
	// Name is a field that represents the name of the student
	Name      string
	// EntryDate is a field that represents the entry date of the student
	EntryDate string
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

// RegisterStudent is a method that adds a new student to the bootcamp
func (b *Bootcamp) RegisterStudent(s Student) {
	(*b).Students = append((*b).Students, s)
}
	

func NewBootcamp(org string, lang string) *Bootcamp {
	return &Bootcamp{
		Organization: org,
		Language:     lang,
	}
}

func main() {
	// Create a new bootcamp
	bootcamp := NewBootcamp("MeLi", "Go")
	
	// Create students
	student1 := NewStudent("John", "2020-01-01")
	student2 := NewStudent("Jane", "2020-01-06")
	student3 := NewStudent("Joe", "2020-01-13")

	// Register students to the bootcamp
	bootcamp.RegisterStudent(*student1)
	bootcamp.RegisterStudent(*student2)
	bootcamp.RegisterStudent(*student3)

	// Print bootcamp info
	fmt.Printf("Bootcamp: %v\n", bootcamp)
}