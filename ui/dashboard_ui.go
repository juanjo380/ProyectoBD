package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildDashboardUI(w fyne.Window) fyne.CanvasObject {

	// TÍTULO ESTILIZADO
	titulo := canvas.NewText("📘 Sistema de Gestión – Confecciones", theme.PrimaryColorNamed(theme.ColorBlue))
	titulo.TextSize = 24
	titulo.Alignment = fyne.TextAlignCenter

	// BOTÓN DE CERRAR SESIÓN
	btnLogout := widget.NewButtonWithIcon("Cerrar Sesión", theme.LogoutIcon(), func() {
		w.SetContent(BuildLoginUI(w))
	})

	// MENÚ PRINCIPAL
	menu := container.NewGridWithColumns(2, // Organizar en 2 columnas
		widget.NewButtonWithIcon("👥 Clientes", theme.AccountIcon(), func() {
			w.SetContent(BuildClienteUI(w))
		}),
		widget.NewButtonWithIcon("🏢 Proveedores", theme.DocumentIcon(), func() {
			w.SetContent(BuildProveedorUI(w))
		}),
		widget.NewButtonWithIcon("👔 Uniformes", theme.ColorPaletteIcon(), func() {
			w.SetContent(BuildUniformeUI(w))
		}),
		widget.NewButtonWithIcon("👕 Productos Terminados", theme.StorageIcon(), func() {
			w.SetContent(BuildProductoTUI(w))
		}),
		widget.NewButtonWithIcon("🧵 Materia Prima", theme.FileIcon(), func() {
			w.SetContent(BuildMateriaPrimaUI(w))
		}),
		widget.NewButtonWithIcon("🔄 Produce", theme.HistoryIcon(), func() {
			w.SetContent(BuildProduceUI(w))
		}),
		widget.NewButtonWithIcon("📦 Posee", theme.FolderIcon(), func() {
			w.SetContent(BuildPoseeUI(w))
		}),
		widget.NewButtonWithIcon("📋 Pedidos", theme.ContentPasteIcon(), func() {
			w.SetContent(BuildPedidosUI(w))
		}),
		widget.NewButtonWithIcon("🏫 Colegios", theme.HomeIcon(), func() {
			w.SetContent(BuildColegioUI(w))
		}),
		widget.NewButtonWithIcon("🧾 Facturas", theme.FileTextIcon(), func() {
			w.SetContent(BuildFacturaUI(w))
		}),
		widget.NewSeparator(),
		widget.NewButtonWithIcon("👤 Usuarios del Sistema", theme.SettingsIcon(), func() {
			w.SetContent(BuildUsuarioUI(w))
		}),
	)

	// CONTENIDO PRINCIPAL
	contenido := container.NewVBox(
		container.NewCenter(titulo),
		widget.NewSeparator(),
		menu,
		widget.NewSeparator(),
		container.NewCenter(btnLogout),
	)

	// FONDO CON COLOR
	bg := canvas.NewRectangle(theme.BackgroundColor())

	// LAYOUT FINAL
	return container.NewMax(
		bg,
		container.NewPadded(contenido),
	)
}

/*

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
		widget.NewSeparator(),

		widget.NewButton("👥 Clientes", func() {
			w.SetContent(BuildClienteUI(w))
		}),

		widget.NewButton("🏢 Proveedores", func() {
			w.SetContent(BuildProveedorUI(w))
		}),

		widget.NewButton("👔 Uniformes", func() {
			w.SetContent(BuildUniformeUI(w))
		}),

		widget.NewButton("👕 Productos Terminados", func() {
			w.SetContent(BuildProductoTUI(w))
		}),

		widget.NewButton("🧵 Materia Prima", func() {
			w.SetContent(BuildMateriaPrimaUI(w))
		}),

		widget.NewButton("🔄 Produce", func() {
			w.SetContent(BuildProduceUI(w))
		}),

		widget.NewButton("📦 Posee", func() {
			w.SetContent(BuildPoseeUI(w))
		}),

		widget.NewButton("📋 Pedidos", func() {
			w.SetContent(BuildPedidosUI(w))
		}),

		widget.NewButton("🏫 Colegios", func() {
			w.SetContent(BuildColegioUI(w))
		}),

		widget.NewButton("🧾 Facturas", func() {
			w.SetContent(BuildFacturaUI(w))
		}),

		// NUEVO: BOTÓN PARA GESTIÓN DE USUARIOS
		widget.NewSeparator(),
		widget.NewButton("👤 Usuarios del Sistema", func() {
			w.SetContent(BuildUsuarioUI(w))
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
*/
