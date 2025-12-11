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

func BuildClienteUI(w fyne.Window) fyne.CanvasObject {

	clientes, _ := controllers.GetAllClientes()

	data := [][]string{{"Documento", "Nombre", "Teléfono"}}
	for _, c := range clientes {
		data = append(data, []string{
			c.DocID,
			c.Nombre,
			c.Telefono,
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnAdd := widget.NewButton("Registrar Cliente", func() {
		w.SetContent(BuildCrearClienteUI(w))
	})

	btnEdit := widget.NewButton("Editar Cliente", func() {
		openClienteEditDialog(w)
	})

	btnDelete := widget.NewButton("Eliminar Cliente", func() {
		openClienteDeleteDialog(w)
	})

	btnSalir := widget.NewButton("Salir", func() {
		w.SetContent(BuildDashboardUI(w, nil))
	})

	menu := container.NewVBox(btnAdd, btnEdit, btnDelete, btnSalir)

	return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearClienteUI(w fyne.Window) fyne.CanvasObject {

	doc := widget.NewEntry()
	nombre := widget.NewEntry()
	tel := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Documento", Widget: doc},
			{Text: "Nombre", Widget: nombre},
			{Text: "Teléfono", Widget: tel},
		},
		OnSubmit: func() {
			c := models.Cliente{
				DocID:    doc.Text,
				Nombre:   nombre.Text,
				Telefono: tel.Text,
			}

			controllers.InsertCliente(c)
			w.SetContent(BuildClienteUI(w))
		},
		OnCancel: func() {
			w.SetContent(BuildClienteUI(w))
		},
	}

	return container.NewVBox(widget.NewLabel("Registrar Cliente"), form)
}

func openClienteEditDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Editar Cliente",
		"Buscar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "Documento del cliente", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			cliente, err := controllers.GetClienteByID(id.Text)
			if err != nil {
				dialog.ShowInformation("Error", "Cliente no existe", w)
				return
			}

			w.SetContent(BuildEditarClienteUI(w, cliente))
		},
		w,
	)
}

func BuildEditarClienteUI(w fyne.Window, c *models.Cliente) fyne.CanvasObject {

	nombre := widget.NewEntry()
	nombre.SetText(c.Nombre)

	tel := widget.NewEntry()
	tel.SetText(c.Telefono)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Nombre", Widget: nombre},
			{Text: "Teléfono", Widget: tel},
		},
		OnSubmit: func() {
			c.Nombre = nombre.Text
			c.Telefono = tel.Text

			controllers.UpdateCliente(*c)
			w.SetContent(BuildClienteUI(w))
		},
		OnCancel: func() {
			w.SetContent(BuildClienteUI(w))
		},
	}

	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Editando Cliente: %s", c.DocID)),
		form,
	)
}

func openClienteDeleteDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Eliminar Cliente",
		"Eliminar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "Documento del cliente", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			controllers.DeleteCliente(id.Text)
			w.SetContent(BuildClienteUI(w))
		},
		w)
}
