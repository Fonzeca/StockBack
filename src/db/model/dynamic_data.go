package model

type ProductView struct {
	Id           int32  `json:"id"`
	Nombre       string `json:"nombre"`
	IdContenedor *int32 `json:"id_contenedor,omitempty"`
	Cantidad     int16  `json:"cantidad"`
}

type ContainerView struct {
	Id     int32  `json:"id"`
	Nombre string `json:"nombre"`
}
