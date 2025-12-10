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

func BuildPedidosUI(w fyne.Window) fyne.CanvasObject {

    pedidos, _ := controllers.GetAllPedidos()

    data := [][]string{{"ID", "Artículo", "Fecha Encargo", "Entrega", "Cliente", "Factura", "Abono"}}
    for _, p := range pedidos {
        entrega := "Pendiente"
        if p.FechaEntrega != nil {
            entrega = *p.FechaEntrega
        }

        data = append(data, []string{
            p.IDPedido,
            p.Articulo,
            p.FechaEncargo,
            entrega,
            p.DocIDCliente,
            p.IDFactura,
            fmt.Sprint(p.Abono),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 7 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Pedido", func() {
        w.SetContent(BuildCrearPedidoUI(w))
    })

    btnEdit := widget.NewButton("Editar Pedido", func() {
        BuildEditarPedidoDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Pedido", func() {
        BuildEliminarPedidoDialog(w)
    })

    btnEntregar := widget.NewButton("Marcar como Entregado", func() {
        BuildMarcarEntregadoDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete, btnEntregar)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearPedidoUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    articulo := widget.NewEntry()
    anotaciones := widget.NewMultiLineEntry()
    fechaEnc := widget.NewEntry()
    fechaEntrega := widget.NewEntry()
    abono := widget.NewEntry()
    cliente := widget.NewEntry()
    factura := widget.NewEntry()

    fechaEntrega.SetPlaceHolder("opcional")

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Pedido", Widget: id},
            {Text: "Artículo", Widget: articulo},
            {Text: "Anotaciones", Widget: anotaciones},
            {Text: "Fecha Encargo", Widget: fechaEnc},
            {Text: "Fecha Entrega", Widget: fechaEntrega},
            {Text: "Abono", Widget: abono},
            {Text: "Documento Cliente", Widget: cliente},
            {Text: "ID Factura", Widget: factura},
        },
        OnSubmit: func() {
            abonoInt, _ := strconv.Atoi(abono.Text)

            var fEntrega *string
            if fechaEntrega.Text != "" {
                fEntrega = &fechaEntrega.Text
            }

            p := models.Pedido{
                IDPedido:     id.Text,
                Articulo:     articulo.Text,
                Anotaciones:  anotaciones.Text,
                FechaEncargo: fechaEnc.Text,
                FechaEntrega: fEntrega,
                Abono:        abonoInt,
                DocIDCliente: cliente.Text,
                IDFactura:    factura.Text,
            }

            controllers.InsertPedido(p)
            w.SetContent(BuildPedidosUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildPedidosUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Pedido"),
        form,
    )
}

func BuildEditarPedidoDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del pedido a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                pedido, err := controllers.GetPedidoByID(id.Text)
                if err != nil {
                    id.SetText("No existe")
                    return
                }
                w.SetContent(BuildEditarPedidoUI(w, pedido))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarPedidoUI(w fyne.Window, p *models.Pedido) fyne.CanvasObject {

    articulo := widget.NewEntry()
    articulo.SetText(p.Articulo)

    anot := widget.NewMultiLineEntry()
    anot.SetText(p.Anotaciones)

    fechaEncargo := widget.NewEntry()
    fechaEncargo.SetText(p.FechaEncargo)

    fechaEntrega := widget.NewEntry()
    if p.FechaEntrega != nil {
        fechaEntrega.SetText(*p.FechaEntrega)
    }

    abono := widget.NewEntry()
    abono.SetText(fmt.Sprint(p.Abono))

    cliente := widget.NewEntry()
    cliente.SetText(p.DocIDCliente)

    factura := widget.NewEntry()
    factura.SetText(p.IDFactura)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Artículo", Widget: articulo},
            {Text: "Anotaciones", Widget: anot},
            {Text: "Fecha Encargo", Widget: fechaEncargo},
            {Text: "Fecha Entrega", Widget: fechaEntrega},
            {Text: "Abono", Widget: abono},
            {Text: "Cliente", Widget: cliente},
            {Text: "Factura", Widget: factura},
        },
        OnSubmit: func() {
            abonoInt, _ := strconv.Atoi(abono.Text)

            var fEntrega *string
            if fechaEntrega.Text != "" {
                fEntrega = &fechaEntrega.Text
            } else {
                fEntrega = nil
            }

            p.Articulo = articulo.Text
            p.Anotaciones = anot.Text
            p.FechaEncargo = fechaEncargo.Text
            p.FechaEntrega = fEntrega
            p.Abono = abonoInt
            p.DocIDCliente = cliente.Text
            p.IDFactura = factura.Text

            controllers.UpdatePedido(*p)
            w.SetContent(BuildPedidosUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildPedidosUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Pedido: %s", p.IDPedido)),
        form,
    )
}

func BuildEliminarPedidoDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del pedido a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                controllers.DeletePedido(id.Text)
                w.SetContent(BuildPedidosUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildMarcarEntregadoDialog(w fyne.Window) {
    id := widget.NewEntry()
    fecha := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("Marcar Pedido como Entregado"),
            widget.NewLabel("ID Pedido:"),
            id,
            widget.NewLabel("Fecha de Entrega (YYYY-MM-DD):"),
            fecha,
            widget.NewButton("Confirmar", func() {
                controllers.MarcarPedidoEntregado(id.Text, fecha.Text)
                w.SetContent(BuildPedidosUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
