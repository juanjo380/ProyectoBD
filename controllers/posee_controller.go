package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
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
        INSERT INTO posee (id_posee, id_producto_t, id_pedido)
        VALUES ($1, $2, $3, $4)
    `, p.IDPosee, p.IDProductoT, p.IDPedido)

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
        WHERE id_posee=$4
    `, p.IDProductoT, p.IDPedido, p.IDPosee)

	return err
}

func DeletePosee(id int) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM posee WHERE id_posee=$1", id)
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
		return nil, err
	}

	return &p, nil
}
