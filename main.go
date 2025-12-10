package main

import (
    "ProyectoBD/ui"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {

    a := app.New()
    w := a.NewWindow("Sistema de Confecciones – Proyecto Final")

    // Menú principal
    btnClientes := widget.NewButton("Clientes", func() {
        w.SetContent(ui.BuildClienteUI(w))
    })

    btnProveedores := widget.NewButton("Proveedores", func() {
        w.SetContent(ui.BuildProveedorUI(w))
    })

    btnUniformes := widget.NewButton("Uniformes", func() {
        w.SetContent(ui.BuildUniformeUI(w))
    })

    btnProductoT := widget.NewButton("Productos Terminados", func() {
        w.SetContent(ui.BuildProductoTUI(w))
    })

    btnMateriaPrima := widget.NewButton("Materia Prima", func() {
        w.SetContent(ui.BuildMateriasPrimasUI(w))
    })

    btnProduce := widget.NewButton("Produce", func() {
        w.SetContent(ui.BuildProduceUI(w))
    })

    btnPosee := widget.NewButton("Posee", func() {
        w.SetContent(ui.BuildPoseeUI(w))
    })

    btnPedido := widget.NewButton("Pedidos", func() {
        w.SetContent(ui.BuildPedidosUI(w))
    })

    btnColegio := widget.NewButton("Colegios", func() {
        w.SetContent(ui.BuildColegiosUI(w))
    })

    btnFactura := widget.NewButton("Facturas", func() {
        w.SetContent(ui.BuildFacturasUI(w))
    })

    menu := container.NewVBox(
        widget.NewLabel("📘 Sistema de Gestión – Confecciones"),
        btnClientes,
        btnProveedores,
        btnUniformes,
        btnProductoT,
        btnMateriaPrima,
        btnProduce,
        btnPosee,
        btnPedido,
        btnColegio,
        btnFactura,
    )

    w.SetContent(menu)
    w.Resize(fyne.NewSize(700, 500))
    w.ShowAndRun()
}


