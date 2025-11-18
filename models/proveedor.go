package models

type Proveedor struct {
    NIT            string
    Nombre         string
    Direccion      string
    Telefono       string
    NombreContacto string
}

func NewProveedor(nit, nombre, dir, tel, contacto string) *Proveedor {
    return &Proveedor{
        NIT:            nit,
        Nombre:         nombre,
        Direccion:      dir,
        Telefono:       tel,
        NombreContacto: contacto,
    }
}
