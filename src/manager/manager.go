package manager

import (
	"Mindia/Stock1/Stock/src/db"
	"Mindia/Stock1/Stock/src/db/model"
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

	product := model.Producto{}
	tx := db.Find(&product)

	products := []model.Product{}
	for _, product := range products {
		products = append(products, model.Product{
			Id:           product.Id,
			Nombre:       product.Nombre,
			IdContenedor: product.IdContenedor,
			Cantidad:     product.Cantidad,
		})
	}

	return products, tx.Error
}
