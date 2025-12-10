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

func BuildProductoTUI(w fyne.Window) fyne.CanvasObject {

    productos, _ := controllers.GetAllProductosT()

    data := [][]string{{"ID", "Descripción", "Talla", "Sexo", "Precio Venta", "Existencias"}}
    for _, p := range productos {
        data = append(data, []string{
            fmt.Sprint(p.IDProductoT),
            p.Descripcion,
            p.Talla,
            p.Sexo,
            fmt.Sprint(p.PrecioVenta),
            fmt.Sprint(p.CantidadExist),
        })
    }

    table := widget.NewTable(
        func() (int, int) { return len(data), 6 },
        func() fyne.CanvasObject { return widget.NewLabel("") },
        func(id widget.TableCellID, cell fyne.CanvasObject) {
            cell.(*widget.Label).SetText(data[id.Row][id.Col])
        },
    )

    btnAdd := widget.NewButton("Registrar Producto", func() {
        w.SetContent(BuildCrearProductoTUI(w))
    })

    btnEdit := widget.NewButton("Editar Producto", func() {
        BuildEditarProductoTDialog(w)
    })

    btnDelete := widget.NewButton("Eliminar Producto", func() {
        BuildEliminarProductoTDialog(w)
    })

    menu := container.NewVBox(btnAdd, btnEdit, btnDelete)

    return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearProductoTUI(w fyne.Window) fyne.CanvasObject {

    id := widget.NewEntry()
    desc := widget.NewEntry()
    talla := widget.NewEntry()
    sexo := widget.NewEntry()
    precio := widget.NewEntry()
    exist := widget.NewEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "ID Producto", Widget: id},
            {Text: "Descripción", Widget: desc},
            {Text: "Talla", Widget: talla},
            {Text: "Sexo", Widget: sexo},
            {Text: "Precio Venta", Widget: precio},
            {Text: "Cantidad en Inventario", Widget: exist},
        },
        OnSubmit: func() {
            idInt, _ := strconv.Atoi(id.Text)
            precioInt, _ := strconv.Atoi(precio.Text)
            existInt, _ := strconv.Atoi(exist.Text)

            p := models.ProductoTerminado{
                IDProductoT:   idInt,
                Descripcion:   desc.Text,
                Talla:         talla.Text,
                Sexo:          sexo.Text,
                PrecioVenta:   precioInt,
                CantidadExist: existInt,
            }

            controllers.InsertProductoT(p)
            w.SetContent(BuildProductoTUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProductoTUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel("Registrar Producto Terminado"),
        form,
    )
}

func BuildEditarProductoTDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del producto a editar:"),
            id,
            widget.NewButton("Continuar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err != nil {
                    id.SetText("ID inválido")
                    return
                }

                p, err := controllers.GetProductoTByID(idInt)
                if err != nil {
                    id.SetText("No existe")
                    return
                }

                w.SetContent(BuildEditarProductoTUI(w, p))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}

func BuildEditarProductoTUI(w fyne.Window, p *models.ProductoTerminado) fyne.CanvasObject {

    desc := widget.NewEntry()
    desc.SetText(p.Descripcion)

    talla := widget.NewEntry()
    talla.SetText(p.Talla)

    sexo := widget.NewEntry()
    sexo.SetText(p.Sexo)

    precio := widget.NewEntry()
    precio.SetText(fmt.Sprint(p.PrecioVenta))

    exist := widget.NewEntry()
    exist.SetText(fmt.Sprint(p.CantidadExist))

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Descripción", Widget: desc},
            {Text: "Talla", Widget: talla},
            {Text: "Sexo", Widget: sexo},
            {Text: "Precio Venta", Widget: precio},
            {Text: "Existencia", Widget: exist},
        },
        OnSubmit: func() {
            precioInt, _ := strconv.Atoi(precio.Text)
            existInt, _ := strconv.Atoi(exist.Text)

            p.Descripcion = desc.Text
            p.Talla = talla.Text
            p.Sexo = sexo.Text
            p.PrecioVenta = precioInt
            p.CantidadExist = existInt

            controllers.UpdateProductoT(*p)
            w.SetContent(BuildProductoTUI(w))
        },
        OnCancel: func() {
            w.SetContent(BuildProductoTUI(w))
        },
    }

    return container.NewVBox(
        widget.NewLabel(fmt.Sprintf("Editando Producto: %d", p.IDProductoT)),
        form,
    )
}

func BuildEliminarProductoTDialog(w fyne.Window) {
    id := widget.NewEntry()

    popup := widget.NewModalPopUp(
        container.NewVBox(
            widget.NewLabel("ID del producto a eliminar:"),
            id,
            widget.NewButton("Eliminar", func() {
                idInt, err := strconv.Atoi(id.Text)
                if err == nil {
                    controllers.DeleteProductoT(idInt)
                }
                w.SetContent(BuildProductoTUI(w))
                popup.Hide()
            }),
            widget.NewButton("Cancelar", popup.Hide),
        ),
        w.Canvas(),
    )

    popup.Show()
}
