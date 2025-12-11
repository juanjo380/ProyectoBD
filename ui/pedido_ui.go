package ui

import (
	"ProyectoBD/controllers"
	"ProyectoBD/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildPedidosUI(w fyne.Window) fyne.CanvasObject {

	pedidos, _ := controllers.GetAllPedidos()

	data := [][]string{{"ID Pedido", "Artículo", "Anotaciones", "Fecha Encargo", "Fecha Entrega", "Abono", "Cliente", "Factura"}}
	for _, p := range pedidos {

		fechaEntrega := ""
		if p.FechaEntrega != nil {
			fechaEntrega = *p.FechaEntrega
		}

		data = append(data, []string{
			p.IDPedido,
			p.Anotaciones,
			p.FechaEncargo,
			fechaEntrega,
			fmt.Sprintf("%d", p.Abono),
			p.DocIDCliente,
			p.IDFactura,
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 8 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnAdd := widget.NewButton("Registrar Pedido", func() {
		openCrearPedidoDialog(w)
	})

	btnEdit := widget.NewButton("Editar Pedido", func() {
		openEditarPedidoIDDialog(w)
	})

	btnDelete := widget.NewButton("Eliminar Pedido", func() {
		openEliminarPedidoDialog(w)
	})

	btnSalir := widget.NewButton("Salir", func() {
		w.SetContent(BuildDashboardUI(w, nil))
	})

	menu := container.NewVBox(btnAdd, btnEdit, btnDelete, btnSalir)

	return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearPedidoDialog(w fyne.Window) {

	id := widget.NewEntry()
	articulo := widget.NewEntry()
	anot := widget.NewEntry()
	fechaEncargo := widget.NewEntry()
	fechaEntrega := widget.NewEntry() // puede ser vacío
	abono := widget.NewEntry()
	cliente := widget.NewEntry()
	factura := widget.NewEntry()

	dialog.ShowForm("Registrar Pedido",
		"Guardar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID Pedido", Widget: id},
			{Text: "Artículo", Widget: articulo},
			{Text: "Anotaciones", Widget: anot},
			{Text: "Fecha Encargo (YYYY-MM-DD)", Widget: fechaEncargo},
			{Text: "Fecha Entrega (opcional)", Widget: fechaEntrega},
			{Text: "Abono", Widget: abono},
			{Text: "Documento Cliente", Widget: cliente},
			{Text: "ID Factura", Widget: factura},
		},
		func(ok bool) {
			if !ok {
				return
			}

			// Fecha entrega como puntero
			var fechaEntregaPtr *string
			if fechaEntrega.Text != "" {
				v := fechaEntrega.Text
				fechaEntregaPtr = &v
			}

			abonoVal := 0
			fmt.Sscan(abono.Text, &abonoVal)

			p := models.Pedido{
				IDPedido:     id.Text,
				Anotaciones:  anot.Text,
				FechaEncargo: fechaEncargo.Text,
				FechaEntrega: fechaEntregaPtr,
				Abono:        abonoVal,
				DocIDCliente: cliente.Text,
				IDFactura:    factura.Text,
			}

			controllers.InsertPedido(p)
			w.SetContent(BuildPedidosUI(w))
		},
		w,
	)
}

func openEditarPedidoIDDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Editar Pedido",
		"Buscar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID Pedido", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			pedido, err := controllers.GetPedidoByID(id.Text)
			if err != nil {
				dialog.ShowInformation("Error", "No existe un pedido con ese ID", w)
				return
			}

			w.SetContent(BuildEditarPedidoUI(w, pedido))
		},
		w,
	)
}

func BuildEditarPedidoUI(w fyne.Window, p *models.Pedido) fyne.CanvasObject {

	anot := widget.NewEntry()
	anot.SetText(p.Anotaciones)

	fechaEncargo := widget.NewEntry()
	fechaEncargo.SetText(p.FechaEncargo)

	fechaEntrega := widget.NewEntry()
	if p.FechaEntrega != nil {
		fechaEntrega.SetText(*p.FechaEntrega)
	}

	abono := widget.NewEntry()
	abono.SetText(fmt.Sprintf("%d", p.Abono))

	cliente := widget.NewEntry()
	cliente.SetText(p.DocIDCliente)

	factura := widget.NewEntry()
	factura.SetText(p.IDFactura)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Anotaciones", Widget: anot},
			{Text: "Fecha Encargo", Widget: fechaEncargo},
			{Text: "Fecha Entrega", Widget: fechaEntrega},
			{Text: "Abono", Widget: abono},
			{Text: "Cliente", Widget: cliente},
			{Text: "Factura", Widget: factura},
		},
		OnSubmit: func() {

			p.Anotaciones = anot.Text
			p.FechaEncargo = fechaEncargo.Text

			// fecha entrega opcional
			if fechaEntrega.Text == "" {
				p.FechaEntrega = nil
			} else {
				v := fechaEntrega.Text
				p.FechaEntrega = &v
			}

			fmt.Sscan(abono.Text, &p.Abono)
			p.DocIDCliente = cliente.Text
			p.IDFactura = factura.Text

			controllers.UpdatePedido(*p)
			w.SetContent(BuildPedidosUI(w))
		},
		OnCancel: func() {
			w.SetContent(BuildPedidosUI(w))
		},
	}

	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Editando Pedido: %s", p.IDPedido)),
		form,
	)
}

func openEliminarPedidoDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Eliminar Pedido",
		"Eliminar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID del Pedido", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			controllers.DeletePedido(id.Text)
			w.SetContent(BuildPedidosUI(w))
		},
		w,
	)
}
