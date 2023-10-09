package main

import "fmt"

func main() {
	orch := &Orchestrator{
		handlers: map[string]func(){
			"handler1": HandlerOne,
			"handler2": HandlerTwo,
			"handler3": HandlerThree,
		},
	}

	orch.RunHandler("handler1")
	orch.RunHandler("handler2")
	orch.RunHandler("handler3") // panic
	orch.RunHandler("handler1")
}

// Video-Speech: Aca tenemos un ejemplo de un orchestrator, que es un tipo que se encarga de ejecutar handlers previamente registrados.
// si alguno de los handlers falla, el programa no se detiene, simplemente no se ejecuta el handler. Recover se encarga de capturar el panic
// y el programa continua su ejecución normalmente.
type Orchestrator struct {
	handlers map[string]func()
}

func (o *Orchestrator) RunHandler(name string) {
	// recover from panic
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
