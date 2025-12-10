package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllFacturas() ([]models.Factura, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query("SELECT id_factura, estado, monto_total FROM facturas")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Factura

    for rows.Next() {
        var f models.Factura
        if err := rows.Scan(&f.IDFactura, &f.Estado, &f.MontoTotal); err != nil {
            return nil, err
        }
        lista = append(lista, f)
    }

    return lista, nil
}

func InsertFactura(f models.Factura) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO facturas (id_factura, estado, monto_total)
        VALUES ($1, $2, $3)
    `, f.IDFactura, f.Estado, f.MontoTotal)

    return err
}

func UpdateFactura(f models.Factura) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE facturas 
        SET estado=$1, monto_total=$2 
        WHERE id_factura=$3
    `, f.Estado, f.MontoTotal, f.IDFactura)

    return err
}

func DeleteFactura(id string) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM facturas WHERE id_factura=$1", id)
    return err
}

func GetFacturaByID(id string) (*models.Factura, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var f models.Factura

    err = conn.QueryRow(`
        SELECT id_factura, estado, monto_total 
        FROM facturas 
        WHERE id_factura=$1
    `, id).Scan(&f.IDFactura, &f.Estado, &f.MontoTotal)

    if err != nil {
        return nil, err
    }

    return &f, nil
}
