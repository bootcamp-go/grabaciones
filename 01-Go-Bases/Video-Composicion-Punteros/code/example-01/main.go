package main

// Video Speech:
// Buenas, en este video vamos a aplicar composición y el uso de métodos con punteros
// previamente vimos como crear un struct y como crear un método para ese struct, donde el receiver es una copia del struct.
// aqui podemos ver un struct llamado Person, con un método llamado Fullname, donde el receiver representa una copia del struct
// y el método retorna un string con el nombre completo de la persona.
type Person struct {
	FirstName string
	LastName string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// que ocurre si queremos modificar el nombre de la persona desde el método?
// para esto podemos usar punteros, donde el receptor pasará a ser un puntero al struct, es decir una referencia a la instancia del struct
// y no una copia del mismo, de esta forma, a través de la dereferencia del puntero, podemos tanto obtener la ultima lectura del struct así como modificarlo
// en este caso, el método recibe un puntero a persona, y modifica el nombre de la persona
func (p *Person) SetFirstName(newName string) {
	(*p).FirstName = newName
}

// Creamos una instancia de Person, y llamamos al método SetFirstName, pasando un string
// y luego imprimimos el nombre completo de la persona

func main() {
	person := Person{FirstName: "John", LastName: "Doe"}
	
	person.SetFirstName("Jane")
	println(person.FullName())
}