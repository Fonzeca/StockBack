package manager

import (
	"Mindia/Stock1/Stock/src/db"
	"Mindia/Stock1/Stock/src/db/model"
	"errors"
	"strconv"
	"time"
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

func (ma *Manager) DeleteProductById(idParam string) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	id, idParseErr := strconv.Atoi(idParam)

	if idParseErr != nil {
		return idParseErr
	}

	producto := model.Producto{
		ID: int32(id),
	}
	tx := db.Delete(&producto)

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

func (ma *Manager) DeleteContainerById(idParam string) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	id, idParseErr := strconv.Atoi(idParam)

	if idParseErr != nil {
		return idParseErr
	}

	contenedor := model.Contenedor{
		ID: int32(id),
	}
	tx := db.Delete(&contenedor)

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

func (ma *Manager) AddProductStockById(Idparam string, amountParam string) (model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ProductView{}, err
	}

	id, idParseErr := strconv.Atoi(Idparam)

	if idParseErr != nil {
		return model.ProductView{}, idParseErr
	}

	amount, amountParseErr := strconv.Atoi(amountParam)

	if amountParseErr != nil {
		return model.ProductView{}, amountParseErr
	}

	producto := model.Producto{
		ID: int32(id),
	}
	tx := db.Find(&producto)

	if tx.Error != nil {
		return model.ProductView{}, tx.Error
	}

	productoUpdated := model.Producto{
		ID:           int32(id),
		Nombre:       producto.Nombre,
		IDContenedor: producto.IDContenedor,
		Cantidad:     producto.Cantidad + int32(amount),
	}
	tx = db.Save(&productoUpdated)

	if tx.Error == nil {
		historyErr := createHistory(productoUpdated.ID, int32(amount), "entrada")
		if err != nil {
			return model.ProductView{}, historyErr
		}
	}

	updatedProduct := model.ProductView{
		Nombre:   productoUpdated.Nombre,
		Cantidad: int16(productoUpdated.Cantidad),
	}

	return updatedProduct, tx.Error
}

func (ma *Manager) RemoveProductStockById(Idparam string, amountParam string) (model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return model.ProductView{}, err
	}

	id, idParseErr := strconv.Atoi(Idparam)

	if idParseErr != nil {
		return model.ProductView{}, idParseErr
	}

	amount, amountParseErr := strconv.Atoi(amountParam)

	if amountParseErr != nil {
		return model.ProductView{}, amountParseErr
	}

	producto := model.Producto{
		ID: int32(id),
	}
	tx := db.Find(&producto)

	if tx.Error != nil {
		return model.ProductView{}, tx.Error
	}

	if int32(amount) > producto.Cantidad {
		return model.ProductView{}, errors.New("no hay stock suficiente")
	}

	productoUpdated := model.Producto{
		ID:           int32(id),
		Nombre:       producto.Nombre,
		IDContenedor: producto.IDContenedor,
		Cantidad:     producto.Cantidad - int32(amount),
	}
	tx = db.Save(&productoUpdated)

	if tx.Error == nil {
		historyErr := createHistory(productoUpdated.ID, int32(amount), "salida")
		if err != nil {
			return model.ProductView{}, historyErr
		}
	}

	updatedProduct := model.ProductView{
		Nombre:   productoUpdated.Nombre,
		Cantidad: int16(productoUpdated.Cantidad),
	}

	return updatedProduct, tx.Error
}

func createHistory(productId int32, amount int32, kind string) error {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return err
	}

	historial := model.Historial{
		IDProducto: productId,
		Fecha:      time.Now(),
		Cantidad:   amount,
		Tipo:       kind,
	}

	tx := db.Save(&historial)

	return tx.Error
}
