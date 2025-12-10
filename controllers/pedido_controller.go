package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
	"database/sql"
)

func GetAllPedidos() ([]models.Pedido, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
        SELECT id_pedido, articulo, anotaciones, fecha_encargo, 
               fecha_entrega, abono, doc_id_cliente, id_factura
        FROM pedidos
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Pedido

	for rows.Next() {
		var p models.Pedido
		var fechaEntrega sql.NullString

		err := rows.Scan(
			&p.IDPedido, &p.Anotaciones,
			&p.FechaEncargo, &fechaEntrega,
			&p.Abono, &p.DocIDCliente, &p.IDFactura,
		)
		if err != nil {
			return nil, err
		}

		if fechaEntrega.Valid {
			p.FechaEntrega = &fechaEntrega.String
		} else {
			p.FechaEntrega = nil
		}

		lista = append(lista, p)
	}

	return lista, nil
}

func InsertPedido(p models.Pedido) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        INSERT INTO pedidos (
            id_pedido, articulo, anotaciones, fecha_encargo, 
            fecha_entrega, abono, doc_id_cliente, id_factura
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, p.IDPedido, p.Anotaciones,
		p.FechaEncargo, p.FechaEntrega,
		p.Abono, p.DocIDCliente, p.IDFactura)

	return err
}

func UpdatePedido(p models.Pedido) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        UPDATE pedidos
        SET articulo=$1, anotaciones=$2, fecha_encargo=$3, 
            fecha_entrega=$4, abono=$5, doc_id_cliente=$6, id_factura=$7
        WHERE id_pedido=$8
    `, p.Anotaciones, p.FechaEncargo,
		p.FechaEntrega, p.Abono, p.DocIDCliente,
		p.IDFactura, p.IDPedido)

	return err
}

func DeletePedido(id string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM pedidos WHERE id_pedido=$1", id)
	return err
}

func GetPedidoByID(id string) (*models.Pedido, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var p models.Pedido
	var fechaEntrega sql.NullString

	err = conn.QueryRow(`
        SELECT id_pedido, articulo, anotaciones, fecha_encargo,
               fecha_entrega, abono, doc_id_cliente, id_factura
        FROM pedidos
        WHERE id_pedido=$1
    `, id).Scan(
		&p.IDPedido, &p.Anotaciones,
		&p.FechaEncargo, &fechaEntrega,
		&p.Abono, &p.DocIDCliente, &p.IDFactura,
	)

	if err != nil {
		return nil, err
	}

	if fechaEntrega.Valid {
		p.FechaEntrega = &fechaEntrega.String
	} else {
		p.FechaEntrega = nil
	}

	return &p, nil
}

func MarcarPedidoEntregado(id string, fecha string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(`
        UPDATE pedidos
        SET fecha_entrega=$1
        WHERE id_pedido=$2
    `, fecha, id)

	return err
}
