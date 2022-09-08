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

func (ma *Manager) GetAllProducts() ([]model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.ProductView{}, err
	}

	productos := []model.Producto{}
	tx := db.Find(&productos)

	products := []model.ProductView{}
	for _, producto := range productos {
		products = append(products, model.ProductView{
			Id:           producto.ID,
			Nombre:       producto.Nombre,
			IdContenedor: producto.IDContenedor,
			Cantidad:     int16(producto.Cantidad),
		})
	}

	return products, tx.Error
}

func (ma *Manager) CreateProduct(productRequest model.Producto) (model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ProductView{}, err
	}

	tx := db.Create(&productRequest)

	product := model.ProductView{
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

func (ma *Manager) ModifyProduct(productRequestBody model.Producto) (model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ProductView{}, err
	}

	tx := db.Save(&productRequestBody)

	product := model.ProductView{
		Id:           productRequestBody.ID,
		IdContenedor: productRequestBody.IDContenedor,
		Nombre:       productRequestBody.Nombre,
		Cantidad:     int16(productRequestBody.Cantidad)}

	return product, tx.Error
}

func (ma *Manager) GetAllContainers() ([]model.ContainerView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.ContainerView{}, err
	}

	contenedores := []model.Contenedor{}
	tx := db.Find(&contenedores)

	containers := []model.ContainerView{}
	for _, contenedor := range contenedores {
		containers = append(containers, model.ContainerView{
			Id:     contenedor.ID,
			Nombre: contenedor.Nombre,
		})
	}

	return containers, tx.Error
}

func (ma *Manager) CreateContainer(containerRequestBody model.Contenedor) (model.ContainerView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ContainerView{}, err
	}

	tx := db.Create(&containerRequestBody)

	container := model.ContainerView{
		Id:     containerRequestBody.ID,
		Nombre: containerRequestBody.Nombre,
	}

	return container, tx.Error
}

func (ma *Manager) DeleteContainerById(param string) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	id, parseErr := strconv.Atoi(param)

	if parseErr != nil {
		return parseErr
	}

	container := model.Contenedor{
		ID: int32(id),
	}
	tx := db.Delete(&container)

	return tx.Error
}

func (ma *Manager) ModifyContainer(containerRequestBody model.Contenedor) (model.ContainerView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ContainerView{}, err
	}

	tx := db.Save(&containerRequestBody)

	container := model.ContainerView{
		Id:     containerRequestBody.ID,
		Nombre: containerRequestBody.Nombre,
	}

	return container, tx.Error
}

func (ma *Manager) GetAllHistorys() ([]model.HistoryView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return []model.HistoryView{}, err
	}

	historiales := []model.Historial{}
	tx := db.Find(&historiales)

	historys := []model.HistoryView{}
	for _, historial := range historiales {
		historys = append(historys, model.HistoryView{
			Id:         historial.ID,
			IdProducto: historial.IDProducto,
			Fecha:      historial.Fecha,
			Cantidad:   int16(historial.Cantidad),
			Tipo:       historial.Tipo,
		})
	}

	return historys, tx.Error
}
