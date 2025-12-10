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

func BuildFacturasUI(w fyne.Window) fyne.CanvasObject {

    facturas, _ := controllers.GetAllFacturas()

    data := [][]string{{"ID", "Estado", "Monto"}}
    for _, f := range facturas {
        data = append(data, []string{
            f.IDFactura, f.Estado, fmt.Sprintf("%d", f.MontoTotal),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 3 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Factura", func() {
        w.SetContent(BuildCrearFacturaUI(w))
    })

    btnEdit := widget.NewButton("Editar Factura", func() {
        BuildEditarFacturaDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Factura", func() {
        BuildEliminarFacturaDialog(w)
    })

    leftMenu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(leftMenu, nil, nil, nil, table)
}

func BuildCrearFacturaUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    estado := widget.NewEntry()
    monto := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Factura", Widget: id},
            {Text: "Estado", Widget: estado},
            {Text: "Monto Total", Widget: monto},
        },
        OnSubmit: func() {
            valor, _ := strconv.Atoi(monto.Text)

            f := models.Factura{
                IDFactura:  id.Text,
                Estado:     estado.Text,
                MontoTotal: valor,
            }

            controllers.InsertFactura(f)
            w.SetContent(BuildFacturasUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildFacturasUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Factura"),
        form,
    )
}

func BuildEditarFacturaDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la factura a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                factura, err := controllers.GetFacturaByID(id.Text)
                if err != nil {
                    id.SetText("No existe.")
                    return
                }
                w.SetContent(BuildEditarFacturaUI(w, factura))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
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
            valor, _ := strconv.Atoi(monto.Text)

            f.Estado = estado.Text
            f.MontoTotal = valor

            controllers.UpdateFactura(*f)
            w.SetContent(BuildFacturasUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildFacturasUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Factura: %s", f.IDFactura)),
        form,
    )
}

func BuildEliminarFacturaDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID de la factura a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                controllers.DeleteFactura(id.Text)
                w.SetContent(BuildFacturasUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}
