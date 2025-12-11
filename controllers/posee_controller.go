package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
	"fmt"
)

func GetAllPosee() ([]models.Posee, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
        SELECT id_posee, id_producto_t, id_pedido
        FROM posee
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Posee

	for rows.Next() {
		var p models.Posee
		err := rows.Scan(&p.IDPosee, &p.IDProductoT, &p.IDPedido)
		if err != nil {
			return nil, err
		}
		lista = append(lista, p)
	}

	return lista, nil
}

func InsertPosee(p models.Posee) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        INSERT INTO posee (id_producto_t, id_pedido)
        VALUES ($1, $2)
    `, p.IDProductoT, p.IDPedido)

	if err != nil {
		fmt.Println("ERROR INSERTANDO POSEE:", err)
	} else {
		fmt.Println("POSEE INSERTADO OK")
	}

	return err
}

func UpdatePosee(p models.Posee) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        UPDATE posee
        SET id_producto_t=$1, id_pedido=$2
        WHERE id_posee=$3
    `, p.IDProductoT, p.IDPedido, p.IDPosee)

	if err != nil {
		fmt.Println("ERROR ACTUALIZANDO POSEE:", err)
	}

	return err
}

func DeletePosee(id int) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM posee WHERE id_posee=$1", id)

	if err != nil {
		fmt.Println("ERROR ELIMINANDO POSEE:", err)
	}

	return err
}

func GetPoseeByID(id int) (*models.Posee, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var p models.Posee

	err = conn.QueryRow(`
        SELECT id_posee, id_producto_t, id_pedido
        FROM posee
        WHERE id_posee=$1
    `, id).Scan(&p.IDPosee, &p.IDProductoT, &p.IDPedido)

	if err != nil {
		fmt.Println("ERROR BUSCANDO POSEE:", err)
		return nil, err
	}

	return &p, nil
}
