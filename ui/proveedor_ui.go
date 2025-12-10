package ui

import (
    "ProyectoBD/controllers"
    "ProyectoBD/models"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func BuildProveedorUI(w fyne.Window) fyne.CanvasObject {

    proveedores, _ := controllers.GetAllProveedores()

    data := [][]string{{"NIT", "Nombre", "Dirección", "Teléfono", "Contacto"}}
    for _, p := range proveedores {
        data = append(data, []string{
            p.NIT,
            p.Nombre,
            p.Direccion,
            p.Telefono,
            p.NombreContacto,
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 5 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Proveedor", func() {
        w.SetContent(BuildCrearProveedorUI(w))
    })

    btnEdit := widget.NewButton("Editar Proveedor", func() {
        BuildEditarProveedorDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Proveedor", func() {
        BuildEliminarProveedorDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearProveedorUI(w fyne.Window) fyne.CanvasObject {

    nit := widget.NewEntry()
    nombre := widget.NewEntry()
    direccion := widget.NewEntry()
    telefono := widget.NewEntry()
    contacto := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "NIT", Widget: nit},
            {Text: "Nombre", Widget: nombre},
            {Text: "Dirección", Widget: direccion},
            {Text: "Teléfono", Widget: telefono},
            {Text: "Contacto", Widget: contacto},
        },
        OnSubmit: func() {

            p := models.Proveedor{
                NIT:            nit.Text,
                Nombre:         nombre.Text,
                Direccion:      direccion.Text,
                Telefono:       telefono.Text,
                NombreContacto: contacto.Text,
            }

            controllers.InsertProveedor(p)
            w.SetContent(BuildProveedorUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProveedorUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Proveedor"),
        form,
    )
}

func BuildEditarProveedorDialog(w fyne.Window) {
    nit := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("NIT del proveedor a editar:"),
            nit,
            widget.NewButton("Continuar", func() {
                p, err := controllers.GetProveedorByNIT(nit.Text)
                if err != nil {
                    nit.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarProveedorUI(w, p))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarProveedorUI(w fyne.Window, p *models.Proveedor) fyne.CanvasObject {

    nombre := widget.NewEntry()
    nombre.SetText(p.Nombre)

    direccion := widget.NewEntry()
    direccion.SetText(p.Direccion)

    telefono := widget.NewEntry()
    telefono.SetText(p.Telefono)

    contacto := widget.NewEntry()
    contacto.SetText(p.NombreContacto)

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Nombre", Widget: nombre},
            {Text: "Dirección", Widget: direccion},
            {Text: "Teléfono", Widget: telefono},
            {Text: "Contacto", Widget: contacto},
        },
        OnSubmit: func() {

            p.Nombre = nombre.Text
            p.Direccion = direccion.Text
            p.Telefono = telefono.Text
            p.NombreContacto = contacto.Text

            controllers.UpdateProveedor(*p)
            w.SetContent(BuildProveedorUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProveedorUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Editando Proveedor: " + p.NIT),
        form,
    )
}

func BuildEliminarProveedorDialog(w fyne.Window) {
    nit := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("NIT del proveedor a eliminar:"),
            nit,
            widget.NewButton("Eliminar", func() {
                controllers.DeleteProveedor(nit.Text)
                w.SetContent(BuildProveedorUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
