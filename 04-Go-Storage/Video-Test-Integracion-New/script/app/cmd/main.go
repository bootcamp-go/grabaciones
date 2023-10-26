package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &application.ConfigAppDefault{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		ConfigMySQL: &mysql.Config{
			User:   os.Getenv("MYSQL_USER"),
			Passwd: os.Getenv("MYSQL_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("MYSQL_ADDRESS"),
			DBName: os.Getenv("MYSQL_DATABASE"),
		},
	}
	app := application.NewApplicationDefault(cfg)
	
	// - setup
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