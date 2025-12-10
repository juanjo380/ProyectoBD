package models

type Uniforme struct {
	IDUniforme       int
	TipoPrenda       string
	Color            string
	TipoTela         string
	IDColegio        int
	LlevaBordado     bool
	TipoBordado      string
	UbicacionBordado string
	EsEstampado      bool
	CuelloBordeColor string
	MangasBordeColor string
}

func NewUniforme(id int, prenda, color, tela string, idCol int, llevaBordado bool, tipoBordado, ubicacionBordado string, esEstampado bool, cuelloBordeColor, mangasBordeColor string) *Uniforme {
	return &Uniforme{
		IDUniforme:       id,
		TipoPrenda:       prenda,
		Color:            color,
		TipoTela:         tela,
		LlevaBordado:     llevaBordado,
		TipoBordado:      tipoBordado,
		UbicacionBordado: ubicacionBordado,
		EsEstampado:      esEstampado,
		CuelloBordeColor: cuelloBordeColor,
		MangasBordeColor: mangasBordeColor,
		IDColegio:        idCol,
	}
}
