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

    rows, err := conn.Query(`
        SELECT doc_id, nombre, telefono
        FROM cliente
    `)
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
        INSERT INTO cliente (doc_id, nombre, telefono)
        VALUES ($1, $2, $3)
    `, c.DocID, c.Nombre, c.Telefono)

    return err
}

func UpdateCliente(c models.Cliente) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE cliente
        SET nombre=$1, telefono=$2
        WHERE doc_id=$3
    `, c.Nombre, c.Telefono, c.DocID)

    return err
}

func DeleteCliente(docID string) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM cliente WHERE doc_id=$1", docID)
    return err
}

func GetClienteByID(docID string) (*models.Cliente, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var c models.Cliente

    err = conn.QueryRow(`
        SELECT doc_id, nombre, telefono
        FROM cliente
        WHERE doc_id=$1
    `, docID).Scan(&c.DocID, &c.Nombre, &c.Telefono)

    if err != nil {
        return nil, err
    }

    return &c, nil
}