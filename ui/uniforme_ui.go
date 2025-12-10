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

func BuildUniformeUI(w fyne.Window) fyne.CanvasObject {

    uniformes, _ := controllers.GetAllUniformes()

    data := [][]string{{"ID", "Prenda", "Color", "Tela", "Colegio"}}
    for _, u := range uniformes {
        data = append(data, []string{
            fmt.Sprint(u.IDUniforme),
            u.TipoPrenda,
            u.Color,
            u.TipoTela,
            fmt.Sprint(u.IDColegio),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 5 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Uniforme", func() {
        w.SetContent(BuildCrearUniformeUI(w))
    })

    btnEdit := widget.NewButton("Editar Uniforme", func() {
        BuildEditarUniformeDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Uniforme", func() {
        BuildEliminarUniformeDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearUniformeUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    prenda := widget.NewEntry()
    color := widget.NewEntry()
    tela := widget.NewEntry()
    colegio := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Uniforme", Widget: id},
            {Text: "Tipo de Prenda", Widget: prenda},
            {Text: "Color", Widget: color},
            {Text: "Tipo de Tela", Widget: tela},
            {Text: "ID Colegio", Widget: colegio},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)
            colInt, _ := strconv.Atoi(colegio.Text)

            u := models.Uniforme{
                IDUniforme: idInt,
                TipoPrenda: prenda.Text,
                Color:      color.Text,
                TipoTela:   tela.Text,
                IDColegio:  colInt,
            }

            controllers.InsertUniforme(u)
            w.SetContent(BuildUniformeUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildUniformeUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Uniforme"),
        form,
    )
}

func BuildEditarUniformeDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del uniforme a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                u, err := controllers.GetUniformeByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarUniformeUI(w, u))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarUniformeUI(w fyne.Window, u *models.Uniforme) fyne.CanvasObject {

    prenda := widget.NewEntry()
    prenda.SetText(u.TipoPrenda)

    color := widget.NewEntry()
    color.SetText(u.Color)

    tela := widget.NewEntry()
    tela.SetText(u.TipoTela)

    colegio := widget.NewEntry()
    colegio.SetText(fmt.Sprint(u.IDColegio))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Tipo de Prenda", Widget: prenda},
            {Text: "Color", Widget: color},
            {Text: "Tipo de Tela", Widget: tela},
            {Text: "ID Colegio", Widget: colegio},
        },
        OnSubmit: func() {
            colInt, _ := strconv.Atoi(colegio.Text)

            u.TipoPrenda = prenda.Text
            u.Color = color.Text
            u.TipoTela = tela.Text
            u.IDColegio = colInt

            controllers.UpdateUniforme(*u)
            w.SetContent(BuildUniformeUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildUniformeUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Uniforme: %d", u.IDUniforme)),
        form,
    )
}

func BuildEliminarUniformeDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del uniforme a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeleteUniforme(idInt)
                }
                w.SetContent(BuildUniformeUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
