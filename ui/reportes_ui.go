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

// Helper: botón de volver al menú de reportes (izquierda, arriba)
func makeBackToReportesButton(w fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("← Volver a Reportes", theme.NavigateBackIcon(), func() {
		w.SetContent(BuildReportesUI(w))
	})
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

	// BOTÓN PARA VOLVER (arriba izquierda)
	btnVolver := makeBackToReportesButton(w)

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
	datos, err := controllers.GetClientesConPedidosPendientes()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	data := [][]string{{"DocID", "Cliente", "ID Pedido", "Producto", "Fecha Encargo"}}
	for _, row := range datos {
		data = append(data, []string{
			row["docid"].(string),
			row["nombre"].(string),
			row["id_pedido"].(string),
			row["producto"].(string),
			row["fecha_encargo"].(string),
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 5 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnVolver := makeBackToReportesButton(w)

	return container.NewBorder(
		container.NewHBox(btnVolver),
		nil, nil, nil,
		table,
	)
}

func BuildReporteInventarioUI(w fyne.Window) fyne.CanvasObject {
	datos, err := controllers.GetInventarioDescontandoEncargados()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	data := [][]string{{"ID", "Descripción", "Talla", "Sexo", "Existencia", "Encargados", "Disponibles"}}
	for _, row := range datos {
		data = append(data, []string{
			strconv.Itoa(row["id_producto_t"].(int)),
			row["descripcion"].(string),
			row["talla"].(string),
			row["sexo"].(string),
			strconv.Itoa(row["existencia"].(int)),
			strconv.Itoa(row["encargados"].(int)),
			strconv.Itoa(row["disponibles"].(int)),
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 7 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnVolver := makeBackToReportesButton(w)

	return container.NewBorder(container.NewHBox(btnVolver), nil, nil, nil, table)
}

func BuildReporteColegiosUI(w fyne.Window) fyne.CanvasObject {
	datos, err := controllers.GetColegiosConUniformes()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	data := [][]string{{"ID Colegio", "Nombre"}}
	for _, row := range datos {
		data = append(data, []string{
			strconv.Itoa(row["id_colegio"].(int)),
			row["nombre"].(string),
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 2 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnVolver := makeBackToReportesButton(w)

	return container.NewBorder(container.NewHBox(btnVolver), nil, nil, nil, table)
}

func BuildReporteVentasUI(w fyne.Window) fyne.CanvasObject {
	total, err := controllers.GetTotalVentas()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	porColegio, err := controllers.GetTotalProductosVendidosPorColegio()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	encabezado := widget.NewLabel("Total de Ventas (facturas pagadas): " + strconv.Itoa(total))

	data := [][]string{{"ID Colegio", "Nombre", "Total"}}
	for _, row := range porColegio {
		data = append(data, []string{
			strconv.Itoa(row["id_colegio"].(int)),
			row["nombre"].(string),
			strconv.Itoa(row["total"].(int)),
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	btnVolver := makeBackToReportesButton(w)

	return container.NewBorder(container.NewVBox(btnVolver, encabezado), nil, nil, nil, table)
}

func BuildReporteUniformesColegioUI(w fyne.Window) fyne.CanvasObject {
	// Selector de colegio
	colegios, err := controllers.GetColegiosConUniformes()
	if err != nil {
		return widget.NewLabel("Error: " + err.Error())
	}

	opciones := []string{}
	idPorNombre := map[string]int{}
	for _, c := range colegios {
		nombre := c["nombre"].(string)
		id := c["id_colegio"].(int)
		opciones = append(opciones, nombre)
		idPorNombre[nombre] = id
	}

	selectColegio := widget.NewSelect(opciones, nil)
	selectColegio.PlaceHolder = "Seleccione un colegio"

	// Tabla dinámica
	data := [][]string{{"ID Uniforme", "Prenda", "Talla", "Sexo", "Color", "Descripción"}}
	table := widget.NewTable(
		func() (int, int) { return len(data), 6 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	selectColegio.OnChanged = func(nombre string) {
		id := idPorNombre[nombre]
		filas, err := controllers.GetUniformesPorColegio(id)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		data = [][]string{{"ID Uniforme", "Prenda", "Talla", "Sexo", "Color", "Descripción"}}
		for _, r := range filas {
			data = append(data, []string{
				strconv.Itoa(r["id_uniforme"].(int)),
				r["prenda"].(string),
				r["talla"].(string),
				r["sexo"].(string),
				r["color"].(string),
				r["descripcion"].(string),
			})
		}
		table.Refresh()
	}

	btnVolver := makeBackToReportesButton(w)

	return container.NewBorder(container.NewHBox(btnVolver, selectColegio), nil, nil, nil, table)
}
