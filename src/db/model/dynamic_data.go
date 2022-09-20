package model

import "time"

type ProductView struct {
	Id           int32  `json:"id,omitempty"`
	Nombre       string `json:"nombre,omitempty"`
	IdContenedor *int32 `json:"id_contenedor,omitempty"`
	Cantidad     int16  `json:"cantidad"`
}

type ProductContainerNameView struct {
	Id               int32  `json:"id"`
	Nombre           string `json:"nombre"`
	ContenedorNombre string `json:"nombre_contenedor"`
	Cantidad         int16  `json:"cantidad"`
}

type ContainerView struct {
	Id        int32  `json:"id"`
	Nombre    string `json:"nombre"`
	Categoria string `json:"categoria"`
	Cantidad  int16  `json:"cantidad"`
}

type HistoryView struct {
	Id               int32     `json:"id"`
	NombreProducto   string    `json:"nombre_producto"`
	NombreContenedor string    `json:"nombre_contenedor"`
	Fecha            time.Time `json:"fecha"`
	Cantidad         int16     `json:"cantidad"`
	Tipo             string    `json:"tipo"`
	ContenedorId     int16     `json:"contenedor_id"`
}
