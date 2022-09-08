package main

import (
	"Mindia/Stock1/Stock/src/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api := service.NewApi()

	e.GET("/getAllProducts", api.GetAllProducts)
	e.POST("/createProduct", api.CreateProduct)
	e.DELETE("/deleteProductById", api.DeleteProductById)
	e.PUT("/modifyProduct", api.ModifyProduct)

	e.Logger.Fatal(e.Start(":4747"))
}
