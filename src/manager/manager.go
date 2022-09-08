package manager

import (
	"Mindia/Stock1/Stock/src/db"
	"Mindia/Stock1/Stock/src/db/model"
	"strconv"
)

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (ma *Manager) GetAllProducts() ([]model.Product, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.Product{}, err
	}

	productos := []model.Producto{}
	tx := db.Find(&productos)

	products := []model.Product{}
	for _, product := range productos {
		products = append(products, model.Product{
			Id:           product.ID,
			Nombre:       product.Nombre,
			IdContenedor: product.IDContenedor,
			Cantidad:     int16(product.Cantidad),
		})
	}

	return products, tx.Error
}

func (ma *Manager) CreateProduct(productRequest model.Producto) (model.Product, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.Product{}, err
	}

	tx := db.Create(&productRequest)

	product := model.Product{
		Id:           productRequest.ID,
		IdContenedor: productRequest.IDContenedor,
		Nombre:       productRequest.Nombre,
		Cantidad:     int16(productRequest.Cantidad),
	}

	return product, tx.Error
}

func (ma *Manager) DeleteProductById(param string) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	id, parseErr := strconv.Atoi(param)

	if parseErr != nil {
		return parseErr
	}

	product := model.Producto{
		ID: int32(id),
	}
	tx := db.Delete(&product)

	return tx.Error
}

func (ma *Manager) ModifyProduct(productRequest model.Producto) (model.Product, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.Product{}, err
	}

	tx := db.Save(&productRequest)

	product := model.Product{
		Id:           productRequest.ID,
		IdContenedor: productRequest.IDContenedor,
		Nombre:       productRequest.Nombre,
		Cantidad:     int16(productRequest.Cantidad)}

	return product, tx.Error
}
