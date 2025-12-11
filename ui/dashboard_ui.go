package ui

import (
	"ProyectoBD/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildDashboardUI(w fyne.Window, usuario *models.Usuario) fyne.CanvasObject {
	// Mostrar información del usuario - MANEJAR NIL
	var userInfo string
	if usuario != nil {
		rolTexto := "Vendedor"
		if usuario.Rol == "administrador" {
			rolTexto = "Administrador"
		}
		userInfo = "👤 " + usuario.NombreCompleto + " (" + rolTexto + ")"
	} else {
		userInfo = "👤 Sesión activa"
	}

	// TÍTULO ESTILIZADO
	titulo := canvas.NewText("📘 Sistema de Gestión – Confecciones", theme.PrimaryColorNamed(theme.ColorBlue))
	titulo.TextSize = 24
	titulo.Alignment = fyne.TextAlignCenter

	// Información del usuario
	infoUsuario := canvas.NewText(userInfo, theme.ForegroundColor())
	infoUsuario.TextSize = 16
	infoUsuario.Alignment = fyne.TextAlignCenter

	// BOTÓN DE CERRAR SESIÓN
	btnLogout := widget.NewButtonWithIcon("Cerrar Sesión", theme.LogoutIcon(), func() {
		w.SetContent(BuildWelcomeUI(w))
		w.Resize(fyne.NewSize(600, 400))
	})

	// MENÚ PRINCIPAL - ORGANIZADO EN 2 COLUMNAS
	menu := container.NewGridWithColumns(2,
		// === GESTIÓN DE VENTAS ===
		widget.NewButtonWithIcon("👥 Clientes", theme.AccountIcon(), func() {
			w.SetContent(BuildClienteUI(w))
		}),
		widget.NewButtonWithIcon("📋 Pedidos", theme.ContentPasteIcon(), func() {
			w.SetContent(BuildPedidosUI(w))
		}),
		widget.NewButtonWithIcon("🧾 Facturas", theme.FileTextIcon(), func() {
			w.SetContent(BuildFacturaUI(w))
		}),

		// === GESTIÓN DE PRODUCTOS ===
		widget.NewButtonWithIcon("👕 Productos Terminados", theme.StorageIcon(), func() {
			w.SetContent(BuildProductoTUI(w))
		}),
		widget.NewButtonWithIcon("🧵 Materia Prima", theme.FileIcon(), func() {
			w.SetContent(BuildMateriaPrimaUI(w))
		}),
		widget.NewButtonWithIcon("🏢 Proveedores", theme.DocumentIcon(), func() {
			w.SetContent(BuildProveedorUI(w))
		}),

		// === GESTIÓN ESCOLAR ===
		widget.NewButtonWithIcon("🏫 Colegios", theme.HomeIcon(), func() {
			w.SetContent(BuildColegioUI(w))
		}),
		widget.NewButtonWithIcon("👔 Uniformes", theme.ColorPaletteIcon(), func() {
			w.SetContent(BuildUniformeUI(w))
		}),

		/*// === RELACIONES ===
		widget.NewButtonWithIcon("🔄 Produce", theme.HistoryIcon(), func() {
			w.SetContent(BuildProduceUI(w))
		}),
		widget.NewButtonWithIcon("📦 Posee", theme.FolderIcon(), func() {
			w.SetContent(BuildPoseeUI(w))
		}),*/

		// === BOTÓN DE REPORTES ===
		widget.NewButtonWithIcon("📊 Reportes", theme.InfoIcon(), func() {
			w.SetContent(BuildReportesUI(w))
		}),

		// === ADMINISTRACIÓN ===
		widget.NewButtonWithIcon("👤 Usuarios", theme.SettingsIcon(), func() {
			w.SetContent(BuildUsuarioUI(w))
		}),
	)

	// CONTENIDO PRINCIPAL
	contenido := container.NewVBox(
		container.NewCenter(titulo),
		container.NewCenter(infoUsuario),
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
