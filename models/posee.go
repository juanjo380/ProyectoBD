package models

type Posee struct {
    IDPosee     int
    IDProductoT int
    IDPedido    string
    Cantidad    int
}

func NewPosee(id, idProd int, idPed string, cant int) *Posee {
    return &Posee{
        IDPosee:     id,
        IDProductoT: idProd,
        IDPedido:    idPed,
        Cantidad:    cant,
    }
}
