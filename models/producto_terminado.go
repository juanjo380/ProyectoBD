package models

type ProductoTerminado struct {
    IDProductoT   int
    Descripcion   string
    Talla         string
    Sexo          string
    PrecioVenta   int
    CantidadExist int
}

func NewProductoTerminado(id int, desc, talla, sexo string, precio, cantidad int) *ProductoTerminado {
    return &ProductoTerminado{
        IDProductoT:   id,
        Descripcion:   desc,
        Talla:         talla,
        Sexo:          sexo,
        PrecioVenta:   precio,
        CantidadExist: cantidad,
    }
}
