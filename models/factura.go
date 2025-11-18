package models

type Factura struct {
    IDFactura  string
    Estado     string
    MontoTotal int
}

func NewFactura(id, estado string, monto int) *Factura {
    return &Factura{
        IDFactura:  id,
        Estado:     estado,
        MontoTotal: monto,
    }
}
