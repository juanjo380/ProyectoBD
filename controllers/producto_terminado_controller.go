package controllers

import (
    "ProyectoBD/db"
    "ProyectoBD/models"
)

func GetAllProductosT() ([]models.ProductoTerminado, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    rows, err := conn.Query(`
        SELECT id_producto_t, descripcion, talla, sexo, precio_venta, cantidad_exist
        FROM producto_terminado
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var lista []models.ProductoTerminado

    for rows.Next() {
        var p models.ProductoTerminado
        err := rows.Scan(&p.IDProductoT, &p.Descripcion, &p.Talla, &p.Sexo, &p.PrecioVenta, &p.CantidadExist)
        if err != nil {
            return nil, err
        }

        lista = append(lista, p)
    }

    return lista, nil
}

func InsertProductoT(p models.ProductoTerminado) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        INSERT INTO producto_terminado (id_producto_t, descripcion, talla, sexo, precio_venta, cantidad_exist)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, p.IDProductoT, p.Descripcion, p.Talla, p.Sexo, p.PrecioVenta, p.CantidadExist)

    return err
}

func UpdateProductoT(p models.ProductoTerminado) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`
        UPDATE producto_terminado
        SET descripcion=$1, talla=$2, sexo=$3, precio_venta=$4, cantidad_exist=$5
        WHERE id_producto_t=$6
    `, p.Descripcion, p.Talla, p.Sexo, p.PrecioVenta, p.CantidadExist, p.IDProductoT)

    return err
}

func DeleteProductoT(id int) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec(`DELETE FROM producto_terminado WHERE id_producto_t=$1`, id)
    return err
}

func GetProductoTByID(id int) (*models.ProductoTerminado, error) {
    conn, err := db.Connect()
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    var p models.ProductoTerminado

    err = conn.QueryRow(`
        SELECT id_producto_t, descripcion, talla, sexo, precio_venta, cantidad_exist
        FROM producto_terminado
        WHERE id_producto_t=$1
    `, id).Scan(&p.IDProductoT, &p.Descripcion, &p.Talla, &p.Sexo, &p.PrecioVenta, &p.CantidadExist)

    if err != nil {
        return nil, err
    }

    return &p, nil
}
