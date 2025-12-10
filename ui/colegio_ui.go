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

func BuildColegioUI(w fyne.Window) fyne.CanvasObject {

    colegios, _ := controllers.GetAllColegios()

    data := [][]string{{"ID", "Nombre", "Teléfono", "Dirección"}}
    for _, c := range colegios {
        data = append(data, []string{
            fmt.Sprintf("%d", c.IDColegio),
            c.Nombre,
            c.Telefono,
            c.Direccion,
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 4 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(cell widget.TableCellID, obj fyne.CanvasObject) {
            obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Colegio", func() {
        openCrearColegioDialog(w)
    })

    btnEdit := widget.NewButton("Editar Colegio", func() {
        openEditarColegioIDDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Colegio", func() {
        openEliminarColegioDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearColegioDialog(w fyne.Window) {

    id := widget.NewEntry()
    nombre := widget.NewEntry()
    tel := widget.NewEntry()
    dir := widget.NewEntry()

    dialog.ShowForm("Registrar Colegio",
        "Guardar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID", Widget: id},
            {Text: "Nombre", Widget: nombre},
            {Text: "Teléfono", Widget: tel},
            {Text: "Dirección", Widget: dir},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            c := models.Colegio{
                IDColegio: idVal,
                Nombre:    nombre.Text,
                Telefono:  tel.Text,
                Direccion: dir.Text,
            }

            controllers.InsertColegio(c)
            w.SetContent(BuildColegioUI(w))
        },
        w,
    )
}

func openEditarColegioIDDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Editar Colegio",
        "Buscar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID del colegio", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            colegio, err := controllers.GetColegioByID(idVal)
            if err != nil {
                dialog.ShowInformation("Error", "No existe un colegio con ese ID", w)
                return
            }

            w.SetContent(BuildEditarColegioUI(w, colegio))
        },
        w,
    )
}

func BuildEditarColegioUI(w fyne.Window, c *models.Colegio) fyne.CanvasObject {

    nombre := widget.NewEntry()
    nombre.SetText(c.Nombre)

    tel := widget.NewEntry()
    tel.SetText(c.Telefono)

    dir := widget.NewEntry()
    dir.SetText(c.Direccion)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Nombre", Widget: nombre},
            {Text: "Teléfono", Widget: tel},
            {Text: "Dirección", Widget: dir},
        },
        OnSubmit: func() {

            c.Nombre = nombre.Text
            c.Telefono = tel.Text
            c.Direccion = dir.Text

            controllers.UpdateColegio(*c)
            w.SetContent(BuildColegioUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildColegioUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Colegio ID: %d", c.IDColegio)),
        form,
    )
}

func openEliminarColegioDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Eliminar Colegio",
        "Eliminar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID del colegio", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            controllers.DeleteColegio(idVal)
            w.SetContent(BuildColegioUI(w))
        },
        w,
    )
}

