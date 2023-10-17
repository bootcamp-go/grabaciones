package main

import (
	"app/internal/application"
	"fmt"
	"os"
)

func main() {
	// env
	// ...

	// application
	// - config
	app := application.NewApplicationDefault(application.Config{
		ServerAddr: os.Getenv("SERVER_ADDR"),
	})
	// - set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}