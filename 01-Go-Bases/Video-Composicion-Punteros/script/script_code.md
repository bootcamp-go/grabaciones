Codigo Composicion - Punteros

____________________________________________________________________________________
Intro

Buenas, en este video vamos ver el uso de métodos con punteros
Previamente vimos como crear un struct y como crear un método para ese struct, donde el receiver (es decir el receptor) es una copia del struct Person.

Aqui podemos ver un struct llamado Person que tiene los campos FirstName y LastName de tipo string
Luego implementa un metodo llamado Fullname, donde el receiver es una copia del struct Person y retorna un string con el nombre completo de la persona.

```go
type Person struct {
	FirstName string
	LastName string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}
```

____________________________________________________________________________________
Codigo

Ahora, que ocurre si queremos modificar el primer nombre de la persona pero no queremos que se modifique desde el campo directamente sino que queremos que se modifique desde un metodo. Esto para encapsular la logica de la modificacion del nombre, la cual incluso podria ser mas compleja que simplemente asignar un valor al campo FirstName.

Para esto, podemos crear un metodo llamado SetFirstName, donde el receiver en vez de ser una copia, pasa a ser un puntero a la instancia del struct Person, es decir, un referencia o direccion en memoria hacia la instancia del struct Person. 

A partir de aca, pediremos un nuevo nombre y lo asignaremos al campo FirstName del struct Person.

para poder cambiar el valor del campo FirstName, debemos usar dereferencia con el operador * para poder acceder y modificar el valor del campo FirstName.

```go
func (p *Person) SetFirstName(newName string) {
	(*p).FirstName = newName
}
```

Finalmente crear una instancia del struct Person y llamar al metodo SetFirstName para cambiar el nombre de la persona. Mostramos el nombre completo de la persona y vemos que efectivamente se cambio el nombre.

```go
func main() {
	person := Person{FirstName: "John", LastName: "Doe"}
	
	person.SetFirstName("Jane")
	println(person.FullName())
}
```

Si yo le saco el puntero, vemos que el nombre no se modifica ya que la modificacion se aplica sobre una copia del struct Person.

____________________________________________________________________________________
Ejemplo 2

explicar brevemente
- mostrar
- ejecutar mostrando la informacion del empleado
- migrar la localizacion y mostrar la nueva informacion del empleado


____________________________________________________________________________________
Cierre

En este video aprendimos a como crear métodos de tipo pointer-receiver, donde el receiver es un puntero y nos permite leer asi como modificar el type, o mejor dicho la instancia del type.

Gracias a esto podremos abrir el abanico de posibilidades para implementar nuevos metodos a nuestros types.