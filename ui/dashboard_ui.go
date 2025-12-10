package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func BuildDashboardUI(w fyne.Window) fyne.CanvasObject {

    // BOTÓN DE CERRAR SESIÓN
    btnLogout := widget.NewButton("Cerrar Sesión", func() {
        w.SetContent(BuildLoginUI(w))
    })

    // MENÚ PRINCIPAL
    menu := container.NewVBox(
        widget.NewLabel("📘 Sistema de Gestión – Confecciones"),

        widget.NewButton("Clientes", func() {
            w.SetContent(BuildClienteUI(w))
        }),

        widget.NewButton("Proveedores", func() {
            w.SetContent(BuildProveedorUI(w))
        }),

        widget.NewButton("Uniformes", func() {
            w.SetContent(BuildUniformeUI(w))
        }),

        widget.NewButton("Productos Terminados", func() {
            w.SetContent(BuildProductoTUI(w))
        }),

        widget.NewButton("Materia Prima", func() {
            w.SetContent(BuildMateriaPrimaUI(w))
        }),

        widget.NewButton("Produce", func() {
            w.SetContent(BuildProduceUI(w))
        }),

        widget.NewButton("Posee", func() {
            w.SetContent(BuildPoseeUI(w))
        }),

        widget.NewButton("Pedidos", func() {
            w.SetContent(BuildPedidosUI(w))
        }),

        widget.NewButton("Colegios", func() {
            w.SetContent(BuildColegioUI(w))
        }),

        widget.NewButton("Facturas", func() {
            w.SetContent(BuildFacturaUI(w))
        }),
    )

    // Layout con el botón de logout arriba
    return container.NewBorder(
        btnLogout, // top
        nil,       // bottom
        nil,       // left
        nil,       // right
        menu,      // center
    )
}
