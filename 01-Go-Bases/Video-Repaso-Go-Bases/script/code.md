________________________________________________________
1 - Intro
En este video vamos a hacer un repaso de los conceptos
que fuimos aprediendo a lo largo de Go Bases.

2 - Proyecto Scafolding
Como pueden ver aca, tenemos un archivo main.go con 2 
structs, una para representar estudiantes y otra para un 
bootcamp.

Student tiene 2 campos, uno es name y entrydate que significa
la fecha de ingreso

Bootcamp tiene 3 campos, uno es organization que es la 
organizacion que dicta el bootcamp, otro es language que es
el lenguaje que se va a enseñar y por ultimo tiene un slice
de estudiantes, que representa aquellos registrados.

El objetivo es crear una instancia de un bootcamp y registrar
nuevos estudiantes, particularmente 3. Para ello primero crearemos
una funcion que se llame NewStudent que retornara una instancia de
Student con los campos que le pasemos por parametro.

```go
func NewStudent(name string, entryDate string) *Student {
    return &Student{
        Name:      name,
        EntryDate: entryDate,
    }
}
```

Ahora crearemos una funcion que se llame NewBootcamp que retornara
una instancia de Bootcamp con los campos que le pasemos por parametro.

```go
func NewBootcamp(organization string, language string) *Bootcamp {
    return &Bootcamp{
        Organization: organization,
        Language:     language,
        Students:     []Student{},
    }
}
```

No incluimos el slice de estudiantes porque los registraremos a través
de un metodo que se llame RegisterStudent que recibe un estudiante y lo
agrega al slice de estudiantes del bootcamp.

```go
func (b *Bootcamp) RegisterStudent(s Student) {
    (*b).Students = append((*b).Students, s)
}
```

Ahora en el main, crearemos una instancia de bootcamp y 3 instancias de
estudiantes, y los registraremos en el bootcamp.

```go
func main() {
	// Create a new bootcamp
	bootcamp := NewBootcamp("MeLi", "Go")
	
	// Create students
	student1 := NewStudent("John", "2020-01-01")
	student2 := NewStudent("Jane", "2020-01-06")
	student3 := NewStudent("Joe", "2020-01-13")

	// Add students to the bootcamp
	bootcamp.AddStudent(*student1)
	bootcamp.AddStudent(*student2)
	bootcamp.AddStudent(*student3)

	// Print bootcamp info
	fmt.Printf("Bootcamp: %v\n", bootcamp)
}
```

Como cierre de este video, en este repaso pudimos ver nuevamente structs y métodos, así como crear instancias de structs y llamar a sus métodos.

Un saludo.