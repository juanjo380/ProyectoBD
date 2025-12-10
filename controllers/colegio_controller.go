package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllColegios() ([]models.Colegio, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query("SELECT id_colegio, nombre, telefono, direccion FROM colegios")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Colegio

    for rows.Next() {
        var c models.Colegio
        if err := rows.Scan(&c.IDColegio, &c.Nombre, &c.Telefono, &c.Direccion); err != nil {
            return nil, err
        }
        lista = append(lista, c)
    }

    return lista, nil
}

func InsertColegio(c models.Colegio) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO colegios (id_colegio, nombre, telefono, direccion)
        VALUES ($1, $2, $3, $4)
    `, c.IDColegio, c.Nombre, c.Telefono, c.Direccion)

    return err
}

func UpdateColegio(c models.Colegio) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE colegios 
        SET nombre=$1, telefono=$2, direccion=$3
        WHERE id_colegio=$4
    `, c.Nombre, c.Telefono, c.Direccion, c.IDColegio)

    return err
}

func DeleteColegio(id int) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM colegios WHERE id_colegio=$1", id)
    return err
}

func GetColegioByID(id int) (*models.Colegio, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var c models.Colegio

    err = conn.QueryRow(`
        SELECT id_colegio, nombre, telefono, direccion 
        FROM colegios 
        WHERE id_colegio=$1
    `, id).Scan(&c.IDColegio, &c.Nombre, &c.Telefono, &c.Direccion)

    if err != nil {
        return nil, err
    }

    return &c, nil
}
