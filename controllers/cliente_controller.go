package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
	"fmt"
)

// Obtener todos los clientes
func GetAllClientes() ([]models.Cliente, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
        SELECT docid, nombre, telefono
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

// Insertar un cliente
func InsertCliente(c models.Cliente) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        INSERT INTO cliente (docid, nombre, telefono)
        VALUES ($1, $2, $3)
    `, c.DocID, c.Nombre, c.Telefono)

	if err != nil {
		fmt.Println("ERROR INSERTANDO CLIENTE:", err)
	} else {
		fmt.Println("INSERT REALIZADO OK")
	}

	return err
}

// Actualizar un cliente
func UpdateCliente(c models.Cliente) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        UPDATE cliente
        SET nombre=$1, telefono=$2
        WHERE docid=$3
    `, c.Nombre, c.Telefono, c.DocID)

	if err != nil {
		fmt.Println("ERROR ACTUALIZANDO CLIENTE:", err)
	}

	return err
}

// Eliminar un cliente
func DeleteCliente(docID string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        DELETE FROM cliente
        WHERE docid=$1
    `, docID)

	if err != nil {
		fmt.Println("ERROR ELIMINANDO CLIENTE:", err)
	}

	return err
}

// Obtener un cliente por ID
func GetClienteByID(docID string) (*models.Cliente, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var c models.Cliente

	err = conn.QueryRow(`
        SELECT docid, nombre, telefono
        FROM cliente
        WHERE docid=$1
    `, docID).Scan(&c.DocID, &c.Nombre, &c.Telefono)

	if err != nil {
		fmt.Println("ERROR BUSCANDO CLIENTE:", err)
		return nil, err
	}

	return &c, nil
}
