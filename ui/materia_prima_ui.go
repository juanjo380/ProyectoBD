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

func BuildMateriasPrimasUI(w fyne.Window) fyne.CanvasObject {

    materias, _ := controllers.GetAllMateriasPrimas()

    data := [][]string{{"ID", "Tipo", "Descripción", "Cantidad", "Unidad", "Proveedor"}}
    for _, mp := range materias {
        data = append(data, []string{
            fmt.Sprint(mp.IDMateriaPrima),
            mp.Tipo,
            mp.Descripcion,
            fmt.Sprint(mp.CantidadExist),
            mp.UnidadMedida,
            mp.NitProveedor,
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 6 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Materia Prima", func() {
        w.SetContent(BuildCrearMateriaPrimaUI(w))
    })

    btnEdit := widget.NewButton("Editar Materia Prima", func() {
        BuildEditarMateriaPrimaDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Materia Prima", func() {
        BuildEliminarMateriaPrimaDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearMateriaPrimaUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    tipo := widget.NewEntry()
    desc := widget.NewEntry()
    cant := widget.NewEntry()
    unidad := widget.NewEntry()
    nit := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Materia Prima", Widget: id},
            {Text: "Tipo", Widget: tipo},
            {Text: "Descripción", Widget: desc},
            {Text: "Cantidad Existente", Widget: cant},
            {Text: "Unidad de Medida", Widget: unidad},
            {Text: "NIT Proveedor", Widget: nit},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)
            cantInt, _ := strconv.Atoi(cant.Text)

            mp := models.MateriaPrima{
                IDMateriaPrima: idInt,
                Tipo:           tipo.Text,
                Descripcion:    desc.Text,
                CantidadExist:  cantInt,
                UnidadMedida:   unidad.Text,
                NitProveedor:   nit.Text,
            }

            controllers.InsertMateriaPrima(mp)
            w.SetContent(BuildMateriasPrimasUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildMateriasPrimasUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Materia Prima"),
        form,
    )
}

func BuildEditarMateriaPrimaDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la materia prima a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                mp, err := controllers.GetMateriaPrimaByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarMateriaPrimaUI(w, mp))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarMateriaPrimaUI(w fyne.Window, mp *models.MateriaPrima) fyne.CanvasObject {

    tipo := widget.NewEntry()
    tipo.SetText(mp.Tipo)

    desc := widget.NewEntry()
    desc.SetText(mp.Descripcion)

    cant := widget.NewEntry()
    cant.SetText(fmt.Sprint(mp.CantidadExist))

    unidad := widget.NewEntry()
    unidad.SetText(mp.UnidadMedida)

    nit := widget.NewEntry()
    nit.SetText(mp.NitProveedor)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Tipo", Widget: tipo},
            {Text: "Descripción", Widget: desc},
            {Text: "Cantidad", Widget: cant},
            {Text: "Unidad", Widget: unidad},
            {Text: "NIT Proveedor", Widget: nit},
        },
        OnSubmit: func() {
            cantInt, _ := strconv.Atoi(cant.Text)

            mp.Tipo = tipo.Text
            mp.Descripcion = desc.Text
            mp.CantidadExist = cantInt
            mp.UnidadMedida = unidad.Text
            mp.NitProveedor = nit.Text

            controllers.UpdateMateriaPrima(*mp)
            w.SetContent(BuildMateriasPrimasUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildMateriasPrimasUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Materia Prima: %d", mp.IDMateriaPrima)),
        form,
    )
}

func BuildEliminarMateriaPrimaDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la materia prima a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeleteMateriaPrima(idInt)
                }
                w.SetContent(BuildMateriasPrimasUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}
