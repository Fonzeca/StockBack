package manager

import (
	"Mindia/Stock1/Stock/src/db"
	"Mindia/Stock1/Stock/src/db/model"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
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

	containerUpdatedErr := updateContainer(db, &productRequest, int(productRequest.Cantidad))
	if containerUpdatedErr != nil {
		return model.ProductView{}, containerUpdatedErr
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

	db.Find(&producto)
	containerUpdatedErr := updateContainer(db, &producto, -int(producto.Cantidad))
	if containerUpdatedErr != nil {
		return containerUpdatedErr
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

	producto := model.Producto{
		ID: productRequestBody.ID,
	}
	db.Find(&producto)

	if producto.Cantidad != productRequestBody.Cantidad {
		return model.ProductView{}, errors.New("no está permitido modificar la cantidad")
	}

	tx := db.Save(&productRequestBody)

	product := model.ProductView{
		Id:           productRequestBody.ID,
		IdContenedor: productRequestBody.IDContenedor,
		Nombre:       productRequestBody.Nombre,
		Cantidad:     int16(productRequestBody.Cantidad)}

	return product, tx.Error
}

func (ma *Manager) GetProductsByContainerId(idParam string) ([]model.ProductView, error) {
	db, close, err := db.ObtenerConexionDb()
	defer close()

	if err != nil {
		return nil, err
	}

	id, idParseErr := strconv.Atoi(idParam)

	if idParseErr != nil {
		return nil, idParseErr
	}

	productos := []model.Producto{}
	tx := db.Where("id_contenedor", id).Find(&productos)

	products := []model.ProductView{}
	for _, producto := range productos {
		products = append(products, model.ProductView{
			Id:       producto.ID,
			Nombre:   producto.Nombre,
			Cantidad: int16(producto.Cantidad),
		})
	}

	return products, tx.Error
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
			Id:        contenedor.ID,
			Nombre:    contenedor.Nombre,
			Categoria: contenedor.Categoria,
			Cantidad:  int16(contenedor.Cantidad),
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
		Id:        containerRequestBody.ID,
		Nombre:    containerRequestBody.Nombre,
		Categoria: containerRequestBody.Categoria,
		Cantidad:  int16(containerRequestBody.Cantidad),
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
		Id:        containerRequestBody.ID,
		Nombre:    containerRequestBody.Nombre,
		Categoria: containerRequestBody.Categoria,
		Cantidad:  int16(containerRequestBody.Cantidad),
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

	updatedProduct, productUpdatedErr := updateProduct(db, &producto, amount)

	if productUpdatedErr != nil {
		return model.ProductView{}, productUpdatedErr
	}

	historyErr := createHistory(producto.ID, int32(amount), "entrada")

	if historyErr != nil {
		return model.ProductView{}, historyErr
	}

	containerUpdatedErr := updateContainer(db, &producto, amount)

	if containerUpdatedErr != nil {
		return model.ProductView{}, containerUpdatedErr
	}

	return updatedProduct, nil
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
	updatedProduct, productUpdatedErr := updateProduct(db, &producto, -amount)

	if productUpdatedErr != nil {
		return model.ProductView{}, productUpdatedErr
	}

	historyErr := createHistory(producto.ID, int32(amount), "salida")

	if historyErr != nil {
		return model.ProductView{}, historyErr
	}

	containerUpdatedErr := updateContainer(db, &producto, -amount)

	if containerUpdatedErr != nil {
		return model.ProductView{}, containerUpdatedErr
	}

	return updatedProduct, nil
}

func updateProduct(db *gorm.DB, producto *model.Producto, amount int) (model.ProductView, error) {
	tx := db.Find(&producto)

	if tx.Error != nil {
		return model.ProductView{}, tx.Error
	}

	productoUpdated := model.Producto{
		ID:           producto.ID,
		Nombre:       producto.Nombre,
		IDContenedor: producto.IDContenedor,
		Cantidad:     producto.Cantidad + int32(amount),
	}

	if amount < 0 && int32(amount) < -producto.Cantidad {
		return model.ProductView{}, errors.New("no hay stock suficiente")
	}

	tx = db.Save(&productoUpdated)

	updatedProduct := model.ProductView{
		Nombre:   productoUpdated.Nombre,
		Cantidad: int16(productoUpdated.Cantidad),
	}

	return updatedProduct, tx.Error
}

func updateContainer(db *gorm.DB, producto *model.Producto, amount int) error {
	contenedor := model.Contenedor{
		ID: *producto.IDContenedor,
	}

	tx := db.Find(&contenedor)

	if tx.Error != nil {
		return tx.Error
	}

	contenedorUpdated := model.Contenedor{
		ID:        contenedor.ID,
		Nombre:    contenedor.Nombre,
		Categoria: contenedor.Categoria,
		Cantidad:  contenedor.Cantidad + int32(amount),
	}
	tx = db.Save(&contenedorUpdated)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
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
