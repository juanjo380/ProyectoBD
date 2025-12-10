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

func BuildProduceUI(w fyne.Window) fyne.CanvasObject {

    produceList, _ := controllers.GetAllProduce()

    data := [][]string{{"ID Produce", "ID Materia Prima", "ID Producto T"}}
    for _, pr := range produceList {
        data = append(data, []string{
            fmt.Sprintf("%d", pr.IDProduce),
            fmt.Sprintf("%d", pr.IDMateriaPrima),
            fmt.Sprintf("%d", pr.IDProductoT),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 3 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(c widget.TableCellID, obj fyne.CanvasObject) {
            obj.(*widget.Label).SetText(data[c.Row][c.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Produce", func() {
        openCrearProduceDialog(w)
    })

    btnEdit := widget.NewButton("Editar Produce", func() {
        openEditarProduceIDDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Produce", func() {
        openEliminarProduceDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearProduceDialog(w fyne.Window) {

    idProduce := widget.NewEntry()
    idMat := widget.NewEntry()
    idProd := widget.NewEntry()

    dialog.ShowForm("Registrar Relación Produce",
        "Guardar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID Produce", Widget: idProduce},
            {Text: "ID Materia Prima", Widget: idMat},
            {Text: "ID Producto Terminado", Widget: idProd},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idProduceVal, _ := strconv.Atoi(idProduce.Text)
            idMatVal, _ := strconv.Atoi(idMat.Text)
            idProdVal, _ := strconv.Atoi(idProd.Text)

            pr := models.Produce{
                IDProduce:      idProduceVal,
                IDMateriaPrima: idMatVal,
                IDProductoT:    idProdVal,
            }

            controllers.InsertProduce(pr)
            w.SetContent(BuildProduceUI(w))
        },
        w,
    )
}

func openEditarProduceIDDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Editar Produce",
        "Buscar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID Produce", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            prod, err := controllers.GetProduceByID(idVal)
            if err != nil {
                dialog.ShowInformation("Error", "No existe un registro con ese ID", w)
                return
            }

            w.SetContent(BuildEditarProduceUI(w, prod))
        },
        w,
    )
}

func BuildEditarProduceUI(w fyne.Window, pr *models.Produce) fyne.CanvasObject {

    idMat := widget.NewEntry()
    idMat.SetText(fmt.Sprintf("%d", pr.IDMateriaPrima))

    idProd := widget.NewEntry()
    idProd.SetText(fmt.Sprintf("%d", pr.IDProductoT))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Materia Prima", Widget: idMat},
            {Text: "ID Producto Terminado", Widget: idProd},
        },
        OnSubmit: func() {

            matVal, _ := strconv.Atoi(idMat.Text)
            prodVal, _ := strconv.Atoi(idProd.Text)

            pr.IDMateriaPrima = matVal
            pr.IDProductoT = prodVal

            controllers.UpdateProduce(*pr)
            w.SetContent(BuildProduceUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProduceUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Produce ID: %d", pr.IDProduce)),
        form,
    )
}

func openEliminarProduceDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Eliminar Produce",
        "Eliminar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID Produce", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            controllers.DeleteProduce(idVal)
            w.SetContent(BuildProduceUI(w))
        },
        w,
    )
}
