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

func BuildMateriaPrimaUI(w fyne.Window) fyne.CanvasObject {

    materias, _ := controllers.GetAllMateriasPrimas()

    data := [][]string{{"ID", "Tipo", "Descripción", "Cantidad", "Unidad", "Proveedor"}}
    for _, m := range materias {
        data = append(data, []string{
            fmt.Sprintf("%d", m.IDMateriaPrima),
            m.Tipo,
            m.Descripcion,
            fmt.Sprintf("%d", m.CantidadExist),
            m.UnidadMedida,
            m.NitProveedor,
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 6 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(cell widget.TableCellID, obj fyne.CanvasObject) {
            obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Materia Prima", func() {
        openCrearMateriaPrimaDialog(w)
    })

    btnEdit := widget.NewButton("Editar Materia Prima", func() {
        openEditarMateriaPrimaIDDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Materia Prima", func() {
        openEliminarMateriaPrimaDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearMateriaPrimaDialog(w fyne.Window) {

    id := widget.NewEntry()
    tipo := widget.NewEntry()
    desc := widget.NewEntry()
    cant := widget.NewEntry()
    unidad := widget.NewEntry()
    nit := widget.NewEntry()

    dialog.ShowForm("Registrar Materia Prima",
        "Guardar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID", Widget: id},
            {Text: "Tipo", Widget: tipo},
            {Text: "Descripción", Widget: desc},
            {Text: "Cantidad", Widget: cant},
            {Text: "Unidad de medida", Widget: unidad},
            {Text: "NIT Proveedor", Widget: nit},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)
            cantVal, _ := strconv.Atoi(cant.Text)

            m := models.MateriaPrima{
                IDMateriaPrima: idVal,
                Tipo:           tipo.Text,
                Descripcion:    desc.Text,
                CantidadExist:  cantVal,
                UnidadMedida:   unidad.Text,
                NitProveedor:   nit.Text,
            }

            controllers.InsertMateriaPrima(m)
            w.SetContent(BuildMateriaPrimaUI(w))
        },
        w,
    )
}

func openEditarMateriaPrimaIDDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Editar Materia Prima",
        "Buscar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID de la materia prima", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            materia, err := controllers.GetMateriaPrimaByID(idVal)
            if err != nil {
                dialog.ShowInformation("Error", "No existe materia prima con ese ID", w)
                return
            }

            w.SetContent(BuildEditarMateriaPrimaUI(w, materia))
        },
        w,
    )
}

func BuildEditarMateriaPrimaUI(w fyne.Window, m *models.MateriaPrima) fyne.CanvasObject {

    tipo := widget.NewEntry()
    tipo.SetText(m.Tipo)

    desc := widget.NewEntry()
    desc.SetText(m.Descripcion)

    cant := widget.NewEntry()
    cant.SetText(fmt.Sprintf("%d", m.CantidadExist))

    unidad := widget.NewEntry()
    unidad.SetText(m.UnidadMedida)

    nit := widget.NewEntry()
    nit.SetText(m.NitProveedor)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Tipo", Widget: tipo},
            {Text: "Descripción", Widget: desc},
            {Text: "Cantidad exist.", Widget: cant},
            {Text: "Unidad de medida", Widget: unidad},
            {Text: "NIT Proveedor", Widget: nit},
        },
        OnSubmit: func() {

            cantVal, _ := strconv.Atoi(cant.Text)

            m.Tipo = tipo.Text
            m.Descripcion = desc.Text
            m.CantidadExist = cantVal
            m.UnidadMedida = unidad.Text
            m.NitProveedor = nit.Text

            controllers.UpdateMateriaPrima(*m)
            w.SetContent(BuildMateriaPrimaUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildMateriaPrimaUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Materia Prima: %d", m.IDMateriaPrima)),
        form,
    )
}

func openEliminarMateriaPrimaDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Eliminar Materia Prima",
        "Eliminar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID a eliminar", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            idVal, _ := strconv.Atoi(id.Text)

            controllers.DeleteMateriaPrima(idVal)
            w.SetContent(BuildMateriaPrimaUI(w))
        },
        w,
    )
}
