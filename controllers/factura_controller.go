package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
	"fmt"
)

func GetAllFacturas() ([]models.Factura, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
        SELECT id_factura, estado, monto_total
        FROM factura
        ORDER BY id_factura
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Factura

	for rows.Next() {
		var f models.Factura
		err := rows.Scan(&f.IDFactura, &f.Estado, &f.MontoTotal)
		if err != nil {
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
        INSERT INTO factura (id_factura, estado, monto_total)
        VALUES ($1, $2, $3)
    `, f.IDFactura, f.Estado, f.MontoTotal)

	if err != nil {
		fmt.Println("ERROR INSERTANDO FACTURA:", err)
	} else {
		fmt.Println("FACTURA INSERTADA OK")
	}

	return err
}

func UpdateFactura(f models.Factura) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        UPDATE factura
        SET estado=$1, monto_total=$2
        WHERE id_factura=$3
    `, f.Estado, f.MontoTotal, f.IDFactura)

	if err != nil {
		fmt.Println("ERROR ACTUALIZANDO FACTURA:", err)
	}

	return err
}

func DeleteFactura(id string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM factura WHERE id_factura=$1", id)

	if err != nil {
		fmt.Println("ERROR ELIMINANDO FACTURA:", err)
	}

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
        FROM factura
        WHERE id_factura=$1
    `, id).Scan(&f.IDFactura, &f.Estado, &f.MontoTotal)

	if err != nil {
		fmt.Println("ERROR BUSCANDO FACTURA:", err)
		return nil, err
	}

	return &f, nil
}
