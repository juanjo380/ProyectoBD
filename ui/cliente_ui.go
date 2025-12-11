package ui

import (
	"ProyectoBD/controllers"
	"ProyectoBD/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/*
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
*/
func BuildClienteUI(w fyne.Window) fyne.CanvasObject {
	// Obtener los datos de los clientes
	clientes, _ := controllers.GetAllClientes()

	// Encabezados de la tabla
	data := [][]string{{"Documento", "Nombre", "Teléfono"}}
	for _, c := range clientes {
		data = append(data, []string{
			c.DocID,
			c.Nombre,
			c.Telefono,
		})
	}

	// Crear la tabla con encabezados estilizados
	table := widget.NewTable(
		func() (int, int) { return len(data), 3 },
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			label.SetText(data[cell.Row][cell.Col])

			// Estilizar encabezados
			if cell.Row == 0 {
				label.TextStyle = fyne.TextStyle{Bold: true}
			}
		},
	)
	table.SetColumnWidth(0, 200) // Ajustar ancho de la columna "Documento"
	table.SetColumnWidth(1, 300) // Ajustar ancho de la columna "Nombre"
	table.SetColumnWidth(2, 200) // Ajustar ancho de la columna "Teléfono"

	// Botones de acción con íconos
	btnAdd := widget.NewButtonWithIcon("Registrar Cliente", theme.ContentAddIcon(), func() {
		w.SetContent(BuildCrearClienteUI(w))
	})

	btnEdit := widget.NewButtonWithIcon("Editar Cliente", theme.DocumentIcon(), func() {
		openClienteEditDialog(w)
	})

	btnDelete := widget.NewButtonWithIcon("Eliminar Cliente", theme.DeleteIcon(), func() {
		openClienteDeleteDialog(w)
	})

	btnSalir := widget.NewButtonWithIcon("Salir", theme.NavigateBackIcon(), func() {
		w.SetContent(BuildDashboardUI(w, nil))
	})

	// Contenedor para los botones
	menu := container.NewHBox(
		btnAdd,
		btnEdit,
		btnDelete,
		btnSalir,
	)

	// Título estilizado
	titulo := widget.NewLabelWithStyle("Gestión de Clientes", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Contenedor principal
	contenido := container.NewVBox(
		titulo,
		widget.NewSeparator(),
		table,
		widget.NewSeparator(),
		menu,
	)

	// Añadir relleno y centrar el contenido
	return container.NewPadded(container.NewVBox(contenido))
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
