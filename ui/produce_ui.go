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

func BuildProduceUI(w fyne.Window) fyne.CanvasObject {

    records, _ := controllers.GetAllProduce()

    data := [][]string{{"ID Produce", "ID Materia Prima", "ID Producto T"}}
    for _, r := range records {
        data = append(data, []string{
            fmt.Sprint(r.IDProduce),
            fmt.Sprint(r.IDMateriaPrima),
            fmt.Sprint(r.IDProductoT),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 3 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Produce", func() {
        w.SetContent(BuildCrearProduceUI(w))
    })

    btnEdit := widget.NewButton("Editar", func() {
        BuildEditarProduceDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar", func() {
        BuildEliminarProduceDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearProduceUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    idMat := widget.NewEntry()
    idProd := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Produce", Widget: id},
            {Text: "ID Materia Prima", Widget: idMat},
            {Text: "ID Producto Terminado", Widget: idProd},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)
            matInt, _ := strconv.Atoi(idMat.Text)
            prodInt, _ := strconv.Atoi(idProd.Text)

            record := models.Produce{
                IDProduce:      idInt,
                IDMateriaPrima: matInt,
                IDProductoT:    prodInt,
            }

            controllers.InsertProduce(record)
            w.SetContent(BuildProduceUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProduceUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Produce"),
        form,
    )
}

func BuildEditarProduceDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de Produce a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                record, err := controllers.GetProduceByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarProduceUI(w, record))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarProduceUI(w fyne.Window, p *models.Produce) fyne.CanvasObject {

    idMat := widget.NewEntry()
    idMat.SetText(fmt.Sprint(p.IDMateriaPrima))

    idProd := widget.NewEntry()
    idProd.SetText(fmt.Sprint(p.IDProductoT))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Materia Prima", Widget: idMat},
            {Text: "ID Producto T", Widget: idProd},
        },
        OnSubmit: func() {
            matInt, _ := strconv.Atoi(idMat.Text)
            prodInt, _ := strconv.Atoi(idProd.Text)

            p.IDMateriaPrima = matInt
            p.IDProductoT = prodInt

            controllers.UpdateProduce(*p)
            w.SetContent(BuildProduceUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProduceUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Produce ID: %d", p.IDProduce)),
        form,
    )
}

func BuildEliminarProduceDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de Produce a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeleteProduce(idInt)
                }
                w.SetContent(BuildProduceUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
