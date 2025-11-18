package models

type Produce struct {
    IDProduce      int
    IDMateriaPrima int
    IDProductoT    int
}

func NewProduce(id, idMat, idProd int) *Produce {
    return &Produce{
        IDProduce:      id,
        IDMateriaPrima: idMat,
        IDProductoT:    idProd,
    }
}
