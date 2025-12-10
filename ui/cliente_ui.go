package ui

import (
    "ProyectoBD/controllers"
    "ProyectoBD/models"
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func BuildClientesUI(w fyne.Window) fyne.CanvasObject {

    clientes, _ := controllers.GetAllClientes()

    data := [][]string{{"Documento", "Nombre", "Teléfono"}}
    for _, c := range clientes {
        data = append(data, []string{
            c.DocID, c.Nombre, c.Telefono,
        })
    }

    // Tabla
    table := widget.NewTable(
        func() (int, int) { return len(data), 3 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    // Botón añadir cliente
    addBtn := widget.NewButton("Agregar Cliente", func() {
        w.SetContent(BuildCrearClienteUI(w))
    })

    // Botón eliminar (abre popup)
    deleteBtn := widget.NewButton("Eliminar Cliente", func() {
        showEliminarClienteDialog(w)
    })

    menu := container.NewVBox(addBtn, deleteBtn)

    return container.NewBorder(
        menu, nil, nil, nil,
        table,
    )
}

func BuildCrearClienteUI(w fyne.Window) fyne.CanvasObject {

    doc := widget.NewEntry()
    nombre := widget.NewEntry()
    telefono := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Documento", Widget: doc},
            {Text: "Nombre", Widget: nombre},
            {Text: "Teléfono", Widget: telefono},
        },
        OnSubmit: func() {
            c := models.Cliente{
                DocID:    doc.Text,
                Nombre:   nombre.Text,
                Telefono: telefono.Text,
            }

            controllers.InsertCliente(c)

            // Vuelve a la lista de clientes
            w.SetContent(BuildClientesUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildClientesUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Cliente"),
        form,
    )
}

func BuildEditarClienteDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("Ingrese el documento del cliente a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                cliente, err := controllers.GetClienteByID(id.Text)
                if err != nil {
                    id.SetText("No existe.")
                    return
                }
                w.SetContent(BuildEditarClienteUI(w, cliente))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarClienteUI(w fyne.Window, c *models.Cliente) fyne.CanvasObject {

    nombre := widget.NewEntry()
    nombre.SetText(c.Nombre)

    telefono := widget.NewEntry()
    telefono.SetText(c.Telefono)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Nombre", Widget: nombre},
            {Text: "Teléfono", Widget: telefono},
        },
        OnSubmit: func() {
            c.Nombre = nombre.Text
            c.Telefono = telefono.Text

            controllers.UpdateCliente(*c)
            w.SetContent(BuildClientesUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildClientesUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Cliente: %s", c.DocID)),
        form,
    )
}

func BuildEliminarClienteDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("Documento del cliente a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                controllers.DeleteCliente(id.Text)
                w.SetContent(BuildClientesUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", func() { popup.Hide() }),
        ),
        w.Canvas(),
    )

    popup.Show()
}
