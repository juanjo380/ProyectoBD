package ui

import (
	"ProyectoBD/controllers"
	"ProyectoBD/models"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildPoseeUI(w fyne.Window) fyne.CanvasObject {

	poseeList, _ := controllers.GetAllPosee()

	data := [][]string{{"ID Posee", "ID Producto T", "ID Pedido"}}
	for _, p := range poseeList {
		data = append(data, []string{
			fmt.Sprintf("%d", p.IDPosee),
			fmt.Sprintf("%d", p.IDProductoT),
			p.IDPedido,
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 4 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnAdd := widget.NewButton("Registrar Relación Posee", func() {
		openCrearPoseeDialog(w)
	})

	btnEdit := widget.NewButton("Editar Relación Posee", func() {
		openEditarPoseeIDDialog(w)
	})

	btnDelete := widget.NewButton("Eliminar Relación Posee", func() {
		openEliminarPoseeDialog(w)
	})

	menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

	return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearPoseeDialog(w fyne.Window) {

	idPosee := widget.NewEntry()
	idProd := widget.NewEntry()
	idPed := widget.NewEntry()

	dialog.ShowForm("Registrar Posee",
		"Guardar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID Posee", Widget: idPosee},
			{Text: "ID Producto Terminado", Widget: idProd},
			{Text: "ID Pedido", Widget: idPed},
		},
		func(ok bool) {
			if !ok {
				return
			}

			idPoseeVal, _ := strconv.Atoi(idPosee.Text)
			idProdVal, _ := strconv.Atoi(idProd.Text)

			p := models.Posee{
				IDPosee:     idPoseeVal,
				IDProductoT: idProdVal,
				IDPedido:    idPed.Text,
			}

			controllers.InsertPosee(p)
			w.SetContent(BuildPoseeUI(w))
		},
		w,
	)
}

func openEditarPoseeIDDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Editar Relación Posee",
		"Buscar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID Posee", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			idVal, _ := strconv.Atoi(id.Text)

			posee, err := controllers.GetPoseeByID(idVal)
			if err != nil {
				dialog.ShowInformation("Error", "No existe registro con ese ID Posee", w)
				return
			}

			w.SetContent(BuildEditarPoseeUI(w, posee))
		},
		w,
	)
}

func BuildEditarPoseeUI(w fyne.Window, p *models.Posee) fyne.CanvasObject {

	idProd := widget.NewEntry()
	idProd.SetText(fmt.Sprintf("%d", p.IDProductoT))

	idPed := widget.NewEntry()
	idPed.SetText(p.IDPedido)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "ID Producto Terminado", Widget: idProd},
			{Text: "ID Pedido", Widget: idPed},
		},
		OnSubmit: func() {

			prodVal, _ := strconv.Atoi(idProd.Text)

			p.IDProductoT = prodVal
			p.IDPedido = idPed.Text

			controllers.UpdatePosee(*p)
			w.SetContent(BuildPoseeUI(w))
		},
		OnCancel: func() {
			w.SetContent(BuildPoseeUI(w))
		},
	}

	return container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Editando Posee ID: %d", p.IDPosee)),
		form,
	)
}

func openEliminarPoseeDialog(w fyne.Window) {

	id := widget.NewEntry()

	dialog.ShowForm("Eliminar Relación Posee",
		"Eliminar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID Posee", Widget: id},
		},
		func(ok bool) {
			if !ok {
				return
			}

			idVal, _ := strconv.Atoi(id.Text)

			controllers.DeletePosee(idVal)
			w.SetContent(BuildPoseeUI(w))
		},
		w,
	)
}
