package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllUniformes() ([]models.Uniforme, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query(`
        SELECT id_uniforme, tipo_prenda, color, tipo_tela, id_colegio
        FROM uniforme
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Uniforme

    for rows.Next() {
        var u models.Uniforme
        err := rows.Scan(&u.IDUniforme, &u.TipoPrenda, &u.Color, &u.TipoTela, &u.IDColegio)
        if err != nil {
            return nil, err
        }
        lista = append(lista, u)
    }

    return lista, nil
}

func InsertUniforme(u models.Uniforme) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO uniforme (id_uniforme, tipo_prenda, color, tipo_tela, id_colegio)
        VALUES ($1, $2, $3, $4, $5)
    `, u.IDUniforme, u.TipoPrenda, u.Color, u.TipoTela, u.IDColegio)

    return err
}

func UpdateUniforme(u models.Uniforme) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE uniforme
        SET tipo_prenda=$1, color=$2, tipo_tela=$3, id_colegio=$4
        WHERE id_uniforme=$5
    `, u.TipoPrenda, u.Color, u.TipoTela, u.IDColegio, u.IDUniforme)

    return err
}

func DeleteUniforme(id int) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`DELETE FROM uniforme WHERE id_uniforme=$1`, id)
    return err
}

func GetUniformeByID(id int) (*models.Uniforme, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var u models.Uniforme

    err = conn.QueryRow(`
        SELECT id_uniforme, tipo_prenda, color, tipo_tela, id_colegio
        FROM uniforme
        WHERE id_uniforme=$1
    `, id).Scan(&u.IDUniforme, &u.TipoPrenda, &u.Color, &u.TipoTela, &u.IDColegio)

    if err != nil {
        return nil, err
    }

    return &u, nil
}
