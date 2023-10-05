package main

import "fmt"

// Video Speech: Ahora aplicando composición, vamos a crear un struct llamado Company en representación de una compañia
// la compania tendra una localización, para eso vamos a crear otro struct
// luego vamos a crear un struct llamado Employee en referencia a un empleado

// location.go
// Location is an entity that represents a location
type Location struct {
	// City is a field that represents the city of the location
	City string
	// Country is a field that represents the country of the location
	Country string
}

// company.go
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

// employee.go
type Employee struct {
	Name     string
	Position string	
	Company	 Company
}

// Information is a method that prints the information of the employee
func (e Employee) Information() {
	fmt.Printf("Employee\n-Name: %s\n-Position: %s\n-Company Name: %s\n-Company Location: %+v\n", e.Name, e.Position, e.Company.Name, e.Company.Location)
}


// Creamos una instancia de Company y Employee, y llamamos al método SetName de Company, pasando un string

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

// Video-Speech: ahora cambiando sacando el puntero del receiver, y dejando el receiver como una copia del struct veamos que ocurre