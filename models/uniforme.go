package models

type Uniforme struct {
    IDUniforme int
    TipoPrenda string
    Color      string
    TipoTela   string
    IDColegio  int
}

func NewUniforme(id int, prenda, color, tela string, idCol int) *Uniforme {
    return &Uniforme{
        IDUniforme: id,
        TipoPrenda: prenda,
        Color:      color,
        TipoTela:   tela,
        IDColegio:  idCol,
    }
}
