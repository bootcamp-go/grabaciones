package main

import "fmt"

func main() {
	// Create a new orchestrator
	or := &Orchestrator{
		handlers: map[string]func(){
			"handler1": HandlerOne,
			"handler2": HandlerTwo,
			"handler3": HandlerThree,
		},
	}

	// Run handler1
	or.RunHandler("handler1")
	or.RunHandler("handler2")
	or.RunHandler("handler3")
	or.RunHandler("handler2")
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
