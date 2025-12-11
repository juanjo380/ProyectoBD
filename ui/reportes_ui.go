package ui

import (
	"ProyectoBD/controllers"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildReportesUI(w fyne.Window) fyne.CanvasObject {
	// TÍTULO
	titulo := canvas.NewText("📊 Reportes del Sistema", theme.PrimaryColorNamed(theme.ColorBlue))
	titulo.TextSize = 20
	titulo.Alignment = fyne.TextAlignCenter

	// BOTÓN PARA VOLVER
	btnVolver := widget.NewButtonWithIcon("← Volver al Dashboard", theme.NavigateBackIcon(), func() {
		// Necesitas pasar el usuario si tu dashboard lo requiere
		// Por ahora usamos nil
		w.SetContent(BuildDashboardUI(w, nil))
	})

	// MENÚ DE REPORTES
	menuReportes := container.NewGridWithColumns(2,
		// REPORTE 1: Pedidos pendientes
		widget.NewButtonWithIcon("📋 Pedidos Pendientes", theme.ListIcon(), func() {
			w.SetContent(BuildReportePedidosPendientesUI(w))
		}),

		// REPORTE 2: Clientes con pedidos
		widget.NewButtonWithIcon("👥 Clientes Pendientes", theme.AccountIcon(), func() {
			w.SetContent(BuildReporteClientesPendientesUI(w))
		}),

		// REPORTE 3: Inventario disponible
		widget.NewButtonWithIcon("📦 Inventario Disponible", theme.StorageIcon(), func() {
			w.SetContent(BuildReporteInventarioUI(w))
		}),

		// REPORTE 4: Colegios con uniformes
		widget.NewButtonWithIcon("🏫 Colegios Uniformes", theme.HomeIcon(), func() {
			w.SetContent(BuildReporteColegiosUI(w))
		}),

		// REPORTE 5: Ventas totales
		widget.NewButtonWithIcon("💰 Ventas Totales", theme.MailAttachmentIcon(), func() {
			w.SetContent(BuildReporteVentasUI(w))
		}),

		// REPORTE 6: Uniformes por colegio
		widget.NewButtonWithIcon("👔 Uniformes por Colegio", theme.ColorPaletteIcon(), func() {
			w.SetContent(BuildReporteUniformesColegioUI(w))
		}),
	)

	// CONTENIDO PRINCIPAL
	contenido := container.NewVBox(
		container.NewCenter(titulo),
		widget.NewSeparator(),
		menuReportes,
		widget.NewSeparator(),
		container.NewCenter(btnVolver),
	)

	// FONDO
	bg := canvas.NewRectangle(theme.BackgroundColor())

	return container.NewMax(
		bg,
		container.NewPadded(contenido),
	)
}

func BuildReportePedidosPendientesUI(w fyne.Window) fyne.CanvasObject {
	// Obtener datos del reporte
	datos, err := controllers.GetPedidosPendientes()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	// Crear tabla
	data := [][]string{
		{"ID Pedido", "Cliente", "Producto", "Fecha Encargo", "Fecha Entrega", "Abono", "Estado"},
	}

	for _, row := range datos {
		data = append(data, []string{
			row["id_pedido"].(string),
			row["cliente"].(string),
			row["producto"].(string),
			row["fecha_encargo"].(string),
			row["fecha_entrega"].(string),
			strconv.Itoa(row["abono"].(int)),
			row["estado"].(string),
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 7 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	// BOTÓN PARA VOLVER
	btnVolver := widget.NewButtonWithIcon("← Volver a Reportes", theme.NavigateBackIcon(), func() {
		w.SetContent(BuildReportesUI(w))
	})

	// BOTÓN PARA EXPORTAR
	btnExportar := widget.NewButtonWithIcon("📥 Exportar", theme.DownloadIcon(), func() {
		dialog.ShowInformation("Exportar", "Reporte exportado correctamente", w)
	})

	// CONTENIDO
	return container.NewBorder(
		container.NewHBox(btnVolver, btnExportar),
		nil, nil, nil,
		table,
	)
}

// Funciones temporales para los otros reportes (las implementaremos después)
func BuildReporteClientesPendientesUI(w fyne.Window) fyne.CanvasObject {
	return widget.NewLabel("Reporte de Clientes Pendientes - Por implementar")
}

func BuildReporteInventarioUI(w fyne.Window) fyne.CanvasObject {
	return widget.NewLabel("Reporte de Inventario - Por implementar")
}

func BuildReporteColegiosUI(w fyne.Window) fyne.CanvasObject {
	return widget.NewLabel("Reporte de Colegios - Por implementar")
}

func BuildReporteVentasUI(w fyne.Window) fyne.CanvasObject {
	return widget.NewLabel("Reporte de Ventas - Por implementar")
}

func BuildReporteUniformesColegioUI(w fyne.Window) fyne.CanvasObject {
	return widget.NewLabel("Reporte de Uniformes por Colegio - Por implementar")
}
