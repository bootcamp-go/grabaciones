# Video - Go Errors

---

## Slide 01 - Introduction
¡Hola a todos! ¿Cómo están?

---

## Slide 02 - Temario

---

## Slide 03 - Error Interface

En el desarrollo de software ocurren diversos tipos de errores. Pero, ¿te preguntaste alguna vez cómo GO nos permite manejar estas situaciones?

En Go, los errores se representan a través de una interfaz nativa llamada error, que contiene un único método llamado Error(). Este método devuelve un mensaje de error en forma de cadena de caracteres.

La ventaja de utilizar una interfaz para representar errores en lugar de simples cadenas de texto es que nos brinda una mayor flexibilidad. Podemos abstraer el proceso de generación de errores mediante diversas implementaciones. Esto significa que podemos manejar errores simples con mensajes de error directos, así como errores más complejos con campos específicos que detallan el problema, o incluso cadenas de errores anidadas.

```go
type error interface {
    Error() string
}
```

---

## Slide 04 - ¿Como creamos un error?

Para crear un error en Go tenemos diversas formas. Tomemos un approach mas manual. Primero debemos contar con algún type que implemente el método Error() string de la interface error

En nuestro caso creamos un type CustomError struct que contiene un campo message de tipo string, y un método Error que devuelve el campo message.
Finalmente creamos una variable de type interface error, donde instanciamos nuestro type CustomError.
```go
type CustomError struct {
    message string
}

func (e *CustomError) Error() string {
    return e.message
}

var ErrCustom error = &CustomError{"Error message"}
```

Go también nos proporciona dos paquetes que simplifican la creación de errores. El primero es `errors`, que nos permite crear un error a partir de una cadena de texto. Internamente, errors.New devuelve una interfaz error, representada por una estructura que implementa el método Error().
```go
err := errors.New("Error message")
```

```go
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

Por otro lado, tenemos el package `fmt`, que nos permite crear un error a partir de un string formateado. De manera similar a errors.New, fmt.Errorf devuelve una interfaz error, representada por un struct que implementa el método error.
```go
err := fmt.Errorf("Error message: %s", "Error")
```

```go
type errorWrap struct {
    msg string
    err error
}
```

A diferencia de errors.New, fmt.Errorf nos permite agregar un error anterior, lo que nos permite crear una cadena de errores. Esto significa que un error puede estar encapsulado por otro, como las capas de una cebolla o las muñecas rusas. Para lograr esto, utilizamos el verbo %w, que nos permite añadir un error anterior al mensaje de error.

```go
func main() {
    errBase := errors.New("Error base")
    errWrap := fmt.Errorf("Error wrap: %w", errBase)
}
```

---

## Slide 05 - ¿Como identificamos un error?

Hasta ahora hemos aprendido cómo generar errores, pero también es importante saber cómo identificarlos.

Una forma sencilla de hacerlo es mediante el operador ==, donde comparamos un error con otro.

```go
var ErrNotFound = errors.New("Not found")

func main() {
    fmt.Println(ErrNotFound == errors.New("Not found")) // false
}
```

---

## Slide 06 - Package errors - Is()

Además de identificar errores de manera directa, el package `errors` nos proporciona funciones útiles como `Is` y `As`, que nos permiten trabajar con la identificación de errores, ya sean encapsulados o no.

- Is: nos permite identificar un error en particular, verificando a lo largo de la cadena de errores si alguno de ellos coincide con el target error en base a su type y valor.

```go
func main() {
    errBase := errors.New("Error base")
    errWrap := fmt.Errorf("Error wrap: %w", errBase)

    // Is (err, target)
    fmt.Println(errors.Is(errWrap, errBase)) // true
}
```
En este caso, el error base quedo encapsulado por el error wrap, y utilizando la funcion Is, se busca dentro de la cadena de errores `errWrap` el error `errBase`, el cual lo encuentra y devuelve true.

---

## Slide 07 - Package errors - As()

- As: nos permite identificar a lo largo de la cadena de errores, si alguno de ellos coincide con nuestro target, un type especifico que implementa la interface error. En caso de encontrarlo, carga nuestro target con el error encontrado.

Es similar a realizar un type assertion para ver la implementacion subyacente de la interface, pero a lo largo de la cadena de errores.

De esta forma podemos identificar y obtener información del type subyacente del error.

```go
type CustomError struct {
    Message string
    Code    int
}
func (e *CustomError) Error() string {
    return e.Message
}

func main() {
    errBase := &CustomError{"Error base", 1}
    errWrap := fmt.Errorf("Error wrap: %w", errBase)

    // As
    var err *CustomError
    if errors.As(errWrap, &err) {
        fmt.Println(err.Code) // 1
    }
}
```

---

## Slide 08 - Manejo de errores

Una gran particularidad de GO es que las funciones pueden ser de multi-retorno, es decir, una función puede devolver múltiples valores.

Como buena práctica, aprovechamos esta característica de GO cuando necesitamos trabajar con una función con potencial de fallo, con el fin de que indique que puede que ocurra un error.

Así, una función devolverá los tipos de datos necesarios y, además, un error. De esta manera, podemos indicar fácilmente si la función ha fallado o no y facilitar la deteción de errores.

```go
func() (result interface{}, err error) {
    // ...
}
```

---

## Slide 09 - Manejo de errores - Ejemplo

Veamos cómo declaramos una función, contemplando que algo puede salir mal

```slide
Generacion

Consideramos casos posibles de fallo
```

```go
var ErrDivideByZero = errors.New("can not divide by zero")

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}
```

Como podemos observar, la funcion divide se encargara de dividir 2 numeros flotantes, donde recibe 2 parametros de tipo float64 y devuelve 2 valores, el resultado de la division y un error.

En caso de que el divisor sea 0, devolvera 0 y un error que previamente instanciamos como ErrDivideByZero.

Caso contrario, devolvera el resultado de la division y nil, indicando que no hubo error.

Importante: a la hora de utilizar esta función, debemos considerar que puede devolver un error, por lo que debemos contemplar esta posibilidad, por ende previo a continuar con el flujo de la función, debemos verificar si el error es nil o no.

```slide
Identificación

A la hora de utilizar esta función, debemos considerar que puede devolver un error
```

```go
func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println(err) // can not divide by zero
        return
    }
    fmt.Println(result)
}
```

Por convención se utiliza la variable err, pero podemos utilizar cualquier nombre de variable.

En caso de que el error sea distinto de nil, imprimimos el error y retornamos de la función, cortando el flujo de la misma. Caso contrario, continuamos con el flujo de codigo normalmente.

No aplicamos a la inversa, es decir, no verificamos si el error es nil y continuamos con el flujo de la función, ya que puede generar codigo anidado y dificultar la lectura del mismo. Siempre verificamos si el error es distinto de nil con el fin de un retorno temprano donde el codigo no debería continuar su flujo, caso contrario continuar tranquilamente.

---

## Slide 10 - Conclusiones

```slide
En este video aprendimos:
- Importancia del manejo de errores durante el desarrollo de software.

- Cuanta más información tengamos sobre el problema que ha ocurrido, más rápido y sencillo será solucionarlo.
    + Rápidos + Fácil
    - Tiempo  - Recursos

- Saber cómo gestionar situaciones de error agrega un valor significativo tanto al programador como al código que crea.
```

En este video aprendimos:
- Comprender cómo manejar errores es fundamental durante el desarrollo de software. Esto garantiza que nuestro programa funcione de acuerdo a nuestras expectativas, incluso cuando ocurre una falla.

- Cuanta más información tengamos sobre el problema que ha ocurrido, más rápido y sencillo será solucionarlo. Por lo tanto, un manejo adecuado de errores nos ahorra tiempo y recursos.

- Saber cómo gestionar situaciones de error agrega un valor significativo tanto al programador como al código que crea. ¡Esperamos que puedan aplicar estos conocimientos en sus próximos proyectos!


---

## Slide 11 - Cierre

Espero que el video les haya sido de utilidad. Un saludo y ¡hasta la próxima!