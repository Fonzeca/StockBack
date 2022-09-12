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
	e.PUT("/addProductStockById", api.AddProductStockById)
	e.PUT("/removeProductStockById", api.RemoveProductStockById)
	e.GET("/getProductsByContainerId", api.GetProductsByContainerId)

	e.GET("/getAllContainers", api.GetAllContainers)
	e.POST("/createContainer", api.CreateContainer)
	e.DELETE("/deleteContainerById", api.DeleteContainerById)
	e.PUT("/modifyContainer", api.ModifyContainer)

	e.GET("/getAllHistorys", api.GetAllHistorys)

	e.Logger.Fatal(e.Start(":4747"))
}
