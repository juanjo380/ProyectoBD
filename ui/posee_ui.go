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

func BuildPoseeUI(w fyne.Window) fyne.CanvasObject {

    posee, _ := controllers.GetAllPosee()

    data := [][]string{{"ID", "Producto T", "ID Pedido", "Cantidad"}}
    for _, p := range posee {
        data = append(data, []string{
            fmt.Sprint(p.IDPosee),
            fmt.Sprint(p.IDProductoT),
            p.IDPedido,
            fmt.Sprint(p.Cantidad),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 4 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Relación (Posee)", func() {
        w.SetContent(BuildCrearPoseeUI(w))
    })

    btnEdit := widget.NewButton("Editar Relación", func() {
        BuildEditarPoseeDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Relación", func() {
        BuildEliminarPoseeDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearPoseeUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    idProd := widget.NewEntry()
    idPed := widget.NewEntry()
    cantidad := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Posee", Widget: id},
            {Text: "ID Producto Terminado", Widget: idProd},
            {Text: "ID Pedido", Widget: idPed},
            {Text: "Cantidad", Widget: cantidad},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)
            prodInt, _ := strconv.Atoi(idProd.Text)
            cantInt, _ := strconv.Atoi(cantidad.Text)

            p := models.Posee{
                IDPosee:     idInt,
                IDProductoT: prodInt,
                IDPedido:    idPed.Text,
                Cantidad:    cantInt,
            }

            controllers.InsertPosee(p)
            w.SetContent(BuildPoseeUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildPoseeUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Relación Posee"),
        form,
    )
}

func BuildEditarPoseeDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la relación a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                p, err := controllers.GetPoseeByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarPoseeUI(w, p))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarPoseeUI(w fyne.Window, p *models.Posee) fyne.CanvasObject {

    idProd := widget.NewEntry()
    idProd.SetText(fmt.Sprint(p.IDProductoT))

    idPed := widget.NewEntry()
    idPed.SetText(p.IDPedido)

    cantidad := widget.NewEntry()
    cantidad.SetText(fmt.Sprint(p.Cantidad))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Producto T", Widget: idProd},
            {Text: "ID Pedido", Widget: idPed},
            {Text: "Cantidad", Widget: cantidad},
        },
        OnSubmit: func() {
            prodInt, _ := strconv.Atoi(idProd.Text)
            cantInt, _ := strconv.Atoi(cantidad.Text)

            p.IDProductoT = prodInt
            p.IDPedido = idPed.Text
            p.Cantidad = cantInt

            controllers.UpdatePosee(*p)
            w.SetContent(BuildPoseeUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildPoseeUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Relación Posee ID: %d", p.IDPosee)),
        form,
    )
}

func BuildEliminarPoseeDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la relación a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeletePosee(idInt)
                }
                w.SetContent(BuildPoseeUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
