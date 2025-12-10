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

func BuildFacturaUI(w fyne.Window) fyne.CanvasObject {

    facturas, _ := controllers.GetAllFacturas()

    data := [][]string{{"ID Factura", "Estado", "Monto"}}
    for _, f := range facturas {
        data = append(data, []string{
            f.IDFactura,
            f.Estado,
            fmt.Sprintf("%d", f.MontoTotal),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 3 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, obj fyne.CanvasObject) {
            obj.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Factura", func() {
        openCrearFacturaDialog(w)
    })

    btnEdit := widget.NewButton("Editar Factura", func() {
        openEditarFacturaIDDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Factura", func() {
        openEliminarFacturaDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func openCrearFacturaDialog(w fyne.Window) {

    id := widget.NewEntry()
    estado := widget.NewEntry()
    monto := widget.NewEntry()

    dialog.ShowForm("Registrar Factura",
        "Guardar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID Factura", Widget: id},
            {Text: "Estado", Widget: estado},
            {Text: "Monto Total", Widget: monto},
        },
        func(ok bool) {
            if !ok {
                return
            }

            montoVal, _ := strconv.Atoi(monto.Text)

            f := models.Factura{
                IDFactura:  id.Text,
                Estado:     estado.Text,
                MontoTotal: montoVal,
            }

            controllers.InsertFactura(f)
            w.SetContent(BuildFacturaUI(w))
        },
        w,
    )
}

func openEditarFacturaIDDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Editar Factura",
        "Buscar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID de la factura", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            factura, err := controllers.GetFacturaByID(id.Text)
            if err != nil {
                dialog.ShowInformation("Error", "No existe una factura con ese ID", w)
                return
            }

            w.SetContent(BuildEditarFacturaUI(w, factura))
        },
        w,
    )
}

func BuildEditarFacturaUI(w fyne.Window, f *models.Factura) fyne.CanvasObject {

    estado := widget.NewEntry()
    estado.SetText(f.Estado)

    monto := widget.NewEntry()
    monto.SetText(fmt.Sprintf("%d", f.MontoTotal))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Estado", Widget: estado},
            {Text: "Monto Total", Widget: monto},
        },
        OnSubmit: func() {

            montoVal, _ := strconv.Atoi(monto.Text)

            f.Estado = estado.Text
            f.MontoTotal = montoVal

            controllers.UpdateFactura(*f)
            w.SetContent(BuildFacturaUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildFacturaUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Factura: %s", f.IDFactura)),
        form,
    )
}

func openEliminarFacturaDialog(w fyne.Window) {

    id := widget.NewEntry()

    dialog.ShowForm("Eliminar Factura",
        "Eliminar",
        "Cancelar",
        []*widget.FormItem{
            {Text: "ID de la factura", Widget: id},
        },
        func(ok bool) {
            if !ok {
                return
            }

            controllers.DeleteFactura(id.Text)
            w.SetContent(BuildFacturaUI(w))
        },
        w,
    )
}
