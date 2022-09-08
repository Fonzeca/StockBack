package model

import "time"

type ProductView struct {
	Id           int32  `json:"id"`
	Nombre       string `json:"nombre,omitempty"`
	IdContenedor *int32 `json:"id_contenedor,omitempty"`
	Cantidad     int16  `json:"cantidad,omitempty"`
}

type ContainerView struct {
	Id     int32  `json:"id"`
	Nombre string `json:"nombre"`
}

type HistoryView struct {
	Id         int32     `json:"id"`
	IdProducto int32     `json:"id_producto"`
	Fecha      time.Time `json:"fecha"`
	Cantidad   int16     `json:"cantidad"`
	Tipo       string    `json:"tipo"`
}
