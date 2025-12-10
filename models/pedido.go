package models

type Pedido struct {
	IDPedido     string
	Anotaciones  string
	FechaEncargo string
	FechaEntrega *string
	Abono        int
	DocIDCliente string
	IDFactura    string
}

func NewPedido(id, anot, fechaE string, fechaEntrega *string, abono int, docID, idFactura string) *Pedido {
	return &Pedido{
		IDPedido:     id,
		Anotaciones:  anot,
		FechaEncargo: fechaE,
		FechaEntrega: fechaEntrega,
		Abono:        abono,
		DocIDCliente: docID,
		IDFactura:    idFactura,
	}
}
