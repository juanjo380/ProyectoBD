package models

type Posee struct {
	IDPosee     int
	IDProductoT int
	IDPedido    string
}

func NewPosee(id, idProd int, idPed string) *Posee {
	return &Posee{
		IDPosee:     id,
		IDProductoT: idProd,
		IDPedido:    idPed,
	}
}
