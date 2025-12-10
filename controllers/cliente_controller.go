package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllClientes() ([]models.Cliente, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query("SELECT doc_id, nombre, telefono FROM clientes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Cliente

    for rows.Next() {
        var c models.Cliente
        err := rows.Scan(&c.DocID, &c.Nombre, &c.Telefono)
        if err != nil {
            return nil, err
        }
        lista = append(lista, c)
    }

    return lista, nil
}

func InsertCliente(c models.Cliente) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO clientes (doc_id, nombre, telefono)
        VALUES ($1, $2, $3)
    `, c.DocID, c.Nombre, c.Telefono)

    return err
}

func DeleteCliente(id string) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM clientes WHERE doc_id=$1", id)
    return err
}

func UpdateCliente(c models.Cliente) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE clientes SET nombre=$1, telefono=$2 WHERE doc_id=$3
    `, c.Nombre, c.Telefono, c.DocID)

    return err
}
