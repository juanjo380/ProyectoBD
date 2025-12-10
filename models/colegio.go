package models

type Colegio struct {
    IDColegio string
    Nombre    string
    Telefono  string
    Direccion string
}

func NewColegio(id int, nombre, tel, dir string) *Colegio {
    return &Colegio{
        IDColegio: id,
        Nombre:    nombre,
        Telefono:  tel,
        Direccion: dir,
    }
}
