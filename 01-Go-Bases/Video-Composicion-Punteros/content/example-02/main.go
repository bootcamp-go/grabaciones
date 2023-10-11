package main

import "fmt"

// Location is an entity that represents a location
type Location struct {
	// City is a field that represents the city of the location
	City string
	// Country is a field that represents the country of the location
	Country string
}

// Company is an entity that represents a company
type Company struct {
	// Name is a field that represents the name of the company
	Name string
	// Location is a field that represents the location of the company
	Location Location
}

// MigrateLocation is a method that migrates the location of the company
func (c *Company) MigrateLocation(newLocation Location) {
	(*c).Location = newLocation
}

// Employee is an entity that represents an employee
type Employee struct {
	// Name is a field that represents the name of the employee
	Name     string
	// Position is a field that represents the position of the employee
	Position string	
	// Company is a field that represents the company of the employee
	Company	 Company
}

// Information is a method that prints the information of the employee
func (e Employee) Information() {
	fmt.Printf("Employee\n-Name: %s\n-Position: %s\n-Company Name: %s\n-Company Location: %+v\n", e.Name, e.Position, e.Company.Name, e.Company.Location)
}


// main.go
func main() {
	// create a new employee
	employee := Employee{
		Name: "John Doe",
		Position: "Software Engineer",
		Company: Company{
			Name: "Labs L.L.C.",
			Location: Location{
				City: "Buenos Aires",
				Country: "Argentina",
			},
		},
	}

	// print the information of the employee
	employee.Information()

	// create a new location
	newLocation := Location{
		City: "New York",
		Country: "United States",
	}
	employee.Company.MigrateLocation(newLocation)

	// print the information of the employee
	employee.Information()
}