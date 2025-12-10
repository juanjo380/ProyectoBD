package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllProveedores() ([]models.Proveedor, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query(`
        SELECT nit, nombre, direccion, telefono, nombre_contacto
        FROM proveedor
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.Proveedor

    for rows.Next() {
        var p models.Proveedor
        err := rows.Scan(&p.NIT, &p.Nombre, &p.Direccion, &p.Telefono, &p.NombreContacto)
        if err != nil {
            return nil, err
        }
        lista = append(lista, p)
    }

    return lista, nil
}

func InsertProveedor(p models.Proveedor) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO proveedor (nit, nombre, direccion, telefono, nombre_contacto)
        VALUES ($1, $2, $3, $4, $5)
    `, p.NIT, p.Nombre, p.Direccion, p.Telefono, p.NombreContacto)

    return err
}

func UpdateProveedor(p models.Proveedor) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE proveedor
        SET nombre=$1, direccion=$2, telefono=$3, nombre_contacto=$4
        WHERE nit=$5
    `, p.Nombre, p.Direccion, p.Telefono, p.NombreContacto, p.NIT)

    return err
}

func DeleteProveedor(nit string) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("DELETE FROM proveedor WHERE nit=$1", nit)
    return err
}

func GetProveedorByNIT(nit string) (*models.Proveedor, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var p models.Proveedor

    err = conn.QueryRow(`
        SELECT nit, nombre, direccion, telefono, nombre_contacto
        FROM proveedor
        WHERE nit=$1
    `, nit).Scan(&p.NIT, &p.Nombre, &p.Direccion, &p.Telefono, &p.NombreContacto)

    if err != nil {
        return nil, err
    }

    return &p, nil
}
