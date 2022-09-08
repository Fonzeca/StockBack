package model

type Product struct {
	Id           int32  `json:"id"`
	Nombre       string `json:"nombre"`
	IdContenedor *int32 `json:"id_contenedor,omitempty"`
	Cantidad     int16  `json:"cantidad"`
}
