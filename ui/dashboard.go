package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func BuildMainUI(w fyne.Window) fyne.CanvasObject {

    title := widget.NewLabel("Sistema de Confecciones - Panel Principal")

    menu := container.NewVBox(
        widget.NewButton("Colegios", func() {
            w.SetContent(BuildColegiosUI(w))
        }),
        widget.NewButton("Clientes", func() {
            w.SetContent(BuildClientesUI(w))
        }),
        widget.NewButton("Productos", func() {
            w.SetContent(BuildProductosUI(w))
        }),
        widget.NewButton("Pedidos", func() {
            w.SetContent(BuildPedidosUI(w))
        }),
        widget.NewButton("Inventario Materias Primas", func() {
            w.SetContent(BuildInventarioUI(w))
        }),
        widget.NewButton("Informes", func() {
            w.SetContent(widget.NewLabel("Aquí irán los informes"))
        }),
        widget.NewButton("Salir", func() {
            w.Close()
        }),
    )

    content := widget.NewLabel("Seleccione un módulo del menú.")

    return container.NewHSplit(menu, container.NewVBox(title, content))
}

