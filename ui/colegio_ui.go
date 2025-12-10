package ui

import (
    "ProyectoBD/controllers"
    "ProyectoBD/models"
    "fmt"
    "strconv"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func BuildColegiosUI(w fyne.Window) fyne.CanvasObject {

    colegios, _ := controllers.GetAllColegios()

    data := [][]string{{"ID", "Nombre", "Teléfono", "Dirección"}}
    for _, c := range colegios {
        data = append(data, []string{
            fmt.Sprint(c.IDColegio), c.Nombre, c.Telefono, c.Direccion,
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 4 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Colegio", func() {
        w.SetContent(BuildCrearColegioUI(w))
    })

    btnEdit := widget.NewButton("Editar Colegio", func() {
        BuildEditarColegioDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Colegio", func() {
        BuildEliminarColegioDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearColegioUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    nombre := widget.NewEntry()
    tel := widget.NewEntry()
    dir := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Colegio", Widget: id},
            {Text: "Nombre", Widget: nombre},
            {Text: "Teléfono", Widget: tel},
            {Text: "Dirección", Widget: dir},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)

            c := models.Colegio{
                IDColegio: idInt,
                Nombre:    nombre.Text,
                Telefono:  tel.Text,
                Direccion: dir.Text,
            }

            controllers.InsertColegio(c)
            w.SetContent(BuildColegiosUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildColegiosUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Colegio"),
        form,
    )
}

func BuildEditarColegioDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del colegio a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                colegio, err := controllers.GetColegioByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarColegioUI(w, colegio))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
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
            w.SetContent(BuildColegiosUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildColegiosUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Colegio: %d", c.IDColegio)),
        form,
    )
}

func BuildEliminarColegioDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del colegio a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeleteColegio(idInt)
                }
                w.SetContent(BuildColegiosUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}
