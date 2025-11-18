package models

type MateriaPrima struct {
    IDMateriaPrima int
    Tipo           string
    Descripcion    string
    CantidadExist  int
    UnidadMedida   string
    NitProveedor   string
}

func NewMateriaPrima(id int, tipo, desc string, cant int, unidad, nit string) *MateriaPrima {
    return &MateriaPrima{
        IDMateriaPrima: id,
        Tipo:           tipo,
        Descripcion:    desc,
        CantidadExist:  cant,
        UnidadMedida:   unidad,
        NitProveedor:   nit,
    }
}
