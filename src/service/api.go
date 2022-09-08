package service

import (
	"Mindia/Stock1/Stock/src/manager"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	manager manager.Manager
}

func NewApi() Api {
	m := manager.NewManager()
	return Api{manager: m}
}

func (api *Api) GetAllProducts(c echo.Context) error {

	products, err := api.manager.GetAllProducts()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}
