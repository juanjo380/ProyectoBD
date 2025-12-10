package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllMateriasPrimas() ([]models.MateriaPrima, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query(`
        SELECT id_materia_prima, tipo, descripcion, cantidad_exist, unidad_medida, nit_proveedor 
        FROM materias_primas
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.MateriaPrima

    for rows.Next() {
        var mp models.MateriaPrima
        if err := rows.Scan(&mp.IDMateriaPrima, &mp.Tipo, &mp.Descripcion, &mp.CantidadExist, &mp.UnidadMedida, &mp.NitProveedor); err != nil {
            return nil, err
        }
        lista = append(lista, mp)
    }

    return lista, nil
}

func InsertMateriaPrima(mp models.MateriaPrima) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO materias_primas (id_materia_prima, tipo, descripcion, cantidad_exist, unidad_medida, nit_proveedor)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, mp.IDMateriaPrima, mp.Tipo, mp.Descripcion, mp.CantidadExist, mp.UnidadMedida, mp.NitProveedor)

    return err
}

func UpdateMateriaPrima(mp models.MateriaPrima) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE materias_primas
        SET tipo=$1, descripcion=$2, cantidad_exist=$3, unidad_medida=$4, nit_proveedor=$5
        WHERE id_materia_prima=$6
    `, mp.Tipo, mp.Descripcion, mp.CantidadExist, mp.UnidadMedida, mp.NitProveedor, mp.IDMateriaPrima)

    return err
}

func DeleteMateriaPrima(id int) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM materias_primas WHERE id_materia_prima=$1", id)
    return err
}

func GetMateriaPrimaByID(id int) (*models.MateriaPrima, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var mp models.MateriaPrima

    err = conn.QueryRow(`
        SELECT id_materia_prima, tipo, descripcion, cantidad_exist, unidad_medida, nit_proveedor 
        FROM materias_primas
        WHERE id_materia_prima=$1
    `, id).Scan(
        &mp.IDMateriaPrima, &mp.Tipo, &mp.Descripcion,
        &mp.CantidadExist, &mp.UnidadMedida, &mp.NitProveedor,
    )

    if err != nil {
        return nil, err
    }

    return &mp, nil
}
