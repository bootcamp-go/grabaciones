package main

import "fmt"

/*
	Notas del Orador:
	- Ahora veremos un ejemplo de como ante un panic, podemos recuperarnos y continuar con la ejecución de nuestro programa.

	- En nuestro ejemplo crearemos un orquestador, que sera un type que se encargara de registrar funciones (en este caso llamadas handlers)
	  en base a una clave y ejecutarlas cuando se lo indiquemos.
	- Crearemos 3 funciones y las registraremos en nuestro orquestador.
	- Una vez registradas, utilizaremos nuestro orquestador para ejecutarlas.

	- Puede ocurrir que alguna de estas funciones falle y genere panic, por ende prevendremos esto y recuperaremos el panic para que nuestro programa no se detenga.
	- Para esto utilizaremos la función recover, que nos permite recuperarnos de un panic y capturar la data indicada en el panic.
*/
func main() {
	or := &Orchestrator{
		handlers: map[string]func(){
			"handler1": HandlerOne,
			"handler2": HandlerTwo,
			"handler3": HandlerThree,
		},
	}

	or.RunHandler("handler1")
	or.RunHandler("handler2")
	or.RunHandler("handler3") // panic
	or.RunHandler("handler1")
}

type Orchestrator struct {
	handlers map[string]func()
}

func (o *Orchestrator) RunHandler(name string) {
	// recover from panic
	// - r contains the data about the generated panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	hd, ok := o.handlers[name]
	if !ok {
		return
	}

	hd()
}

func HandlerOne() {
	println("handler1")
}

func HandlerTwo() {
	println("handler2")
}

func HandlerThree() {
	panic("handler3 failed")
}
