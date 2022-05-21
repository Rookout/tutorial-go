package main

import (
	rook "github.com/Rookout/GoSDK"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/theramis/todo-backend-go-echo/pkg/todos"
	"log"
	"os"
)

func main() {
	rook.Start(rook.RookOptions{})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	todos.RegisterEndPoints(e)

	e.Logger.Fatal(e.Start(":" + port))
}
