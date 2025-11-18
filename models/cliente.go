package models

type Cliente struct {
    DocID    string
    Nombre   string
    Telefono string
}

func NewCliente(id, nombre, tel string) *Cliente {
    return &Cliente{
        DocID:    id,
        Nombre:   nombre,
        Telefono: tel,
    }
}
