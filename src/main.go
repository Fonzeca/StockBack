package main

import (
	"Mindia/Stock1/Stock/src/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api := service.NewApi()

	e.GET("/getAllProducts", api.GetAllProducts)

	e.Logger.Fatal(e.Start(":4747"))
}
