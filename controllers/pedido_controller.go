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

// Obtener pedidos pendientes (fecha_entrega IS NULL)
func GetPedidosPendientes() ([]map[string]interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		SELECT 
			p.id_pedido,
			c.nombre AS cliente,
			pt.descripcion AS producto,
			p.fecha_encargo,
			p.fecha_entrega,
			p.abono
		FROM pedidos p
		JOIN cliente c ON p.doc_id_cliente = c.docid
		JOIN posee po ON p.id_pedido = po.id_pedido
		JOIN producto_terminado pt ON po.id_producto_t = pt.id_producto_t
		WHERE p.fecha_entrega IS NULL
		ORDER BY p.fecha_encargo
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []map[string]interface{}

	for rows.Next() {
		var idPedido, cliente, producto, fechaEncargo string
		var fechaEntrega sql.NullString
		var abono int

		err := rows.Scan(&idPedido, &cliente, &producto, &fechaEncargo, &fechaEntrega, &abono)
		if err != nil {
			return nil, err
		}

		fechaEntregaStr := ""
		if fechaEntrega.Valid {
			fechaEntregaStr = fechaEntrega.String
		}

		resultados = append(resultados, map[string]interface{}{
			"id_pedido":     idPedido,
			"cliente":       cliente,
			"producto":      producto,
			"fecha_encargo": fechaEncargo,
			"fecha_entrega": fechaEntregaStr,
			"abono":         abono,
			"estado":        "Pendiente",
		})
	}

	return resultados, nil
}
