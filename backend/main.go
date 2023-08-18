package main

import (
	"Moonlay/database"
	"Moonlay/pkg/mysql"
	"Moonlay/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseConnection()
	database.MigrationDB()

	router.Routes(e.Group("/api/v1"))

	fmt.Println("Running on Port 5050")

	e.Start("localhost:5050")

}
