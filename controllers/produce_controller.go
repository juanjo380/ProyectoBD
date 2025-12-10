package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllProduce() ([]models.Produce, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query(`
        SELECT id_produce, id_materia_prima, id_producto_t
        FROM produce
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Produce

    for rows.Next() {
        var p models.Produce
        err := rows.Scan(&p.IDProduce, &p.IDMateriaPrima, &p.IDProductoT)
        if err != nil {
            return nil, err
        }
        lista = append(lista, p)
    }
    return lista, nil
}

func InsertProduce(p models.Produce) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO produce (id_produce, id_materia_prima, id_producto_t)
        VALUES ($1, $2, $3)
    `, p.IDProduce, p.IDMateriaPrima, p.IDProductoT)

    return err
}

func UpdateProduce(p models.Produce) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE produce
        SET id_materia_prima=$1, id_producto_t=$2
        WHERE id_produce=$3
    `, p.IDMateriaPrima, p.IDProductoT, p.IDProduce)

    return err
}

func DeleteProduce(id int) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM produce WHERE id_produce=$1", id)
    return err
}

func GetProduceByID(id int) (*models.Produce, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var p models.Produce

    err = conn.QueryRow(`
        SELECT id_produce, id_materia_prima, id_producto_t
        FROM produce
        WHERE id_produce=$1
    `, id).Scan(&p.IDProduce, &p.IDMateriaPrima, &p.IDProductoT)

    if err != nil {
        return nil, err
    }

    return &p, nil
}
