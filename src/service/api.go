package service

import (
	"Mindia/Stock1/Stock/src/db/model"
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

func (api *Api) CreateProduct(c echo.Context) error {
	data := model.Producto{}
	c.Bind(&data)

	product, err := api.manager.CreateProduct(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (api *Api) DeleteProductById(c echo.Context) error {
	val, _ := c.FormParams()
	id := val.Get("id")

	err := api.manager.DeleteProductById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (api *Api) ModifyProduct(c echo.Context) error {
	data := model.Producto{}
	c.Bind(&data)

	product, err := api.manager.ModifyProduct(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (api *Api) GetAllContainers(c echo.Context) error {

	containers, err := api.manager.GetAllContainers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, containers)
}

func (api *Api) CreateContainer(c echo.Context) error {
	data := model.Contenedor{}
	c.Bind(&data)

	container, err := api.manager.CreateContainer(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, container)
}

func (api *Api) DeleteContainerById(c echo.Context) error {
	val, _ := c.FormParams()
	id := val.Get("id")

	err := api.manager.DeleteContainerById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (api *Api) ModifyContainer(c echo.Context) error {
	data := model.Contenedor{}
	c.Bind(&data)

	container, err := api.manager.ModifyContainer(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, container)
}

func (api *Api) GetAllHistorys(c echo.Context) error {

	historys, err := api.manager.GetAllHistorys()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, historys)
}
