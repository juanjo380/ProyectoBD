package ui

import (
	"ProyectoBD/controllers"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Pantalla de bienvenida
func BuildWelcomeUI(w fyne.Window) fyne.CanvasObject {
	// Fondo con color
	bg := canvas.NewRectangle(color.NRGBA{R: 240, G: 245, B: 250, A: 255})

	// Título principal
	titulo := canvas.NewText("Bienvenido/as", color.NRGBA{R: 30, G: 70, B: 120, A: 255})
	titulo.TextSize = 40
	titulo.TextStyle = fyne.TextStyle{Bold: true}

	// Subtítulo
	subtitulo := canvas.NewText("Confecciones Hermanos", color.NRGBA{R: 60, G: 100, B: 150, A: 255})
	subtitulo.TextSize = 24

	// Línea decorativa
	linea := canvas.NewLine(color.NRGBA{R: 180, G: 200, B: 220, A: 255})
	linea.StrokeWidth = 3

	// Botón de login (estilo moderno)

	btnLogin := widget.NewButtonWithIcon("Iniciar Sesión", theme.LoginIcon(), func() {
		w.SetContent(BuildLoginUI(w))
		w.Resize(fyne.NewSize(700, 600))
	})
	btnLogin.Importance = widget.HighImportance
	btnLogin.Resize(fyne.NewSize(200, 50))

	// Botón de salir
	btnSalir := widget.NewButtonWithIcon("Salir", theme.CancelIcon(), func() {
		dialog.ShowConfirm("Salir", "¿Estás seguro de que deseas salir?", func(confirmar bool) {
			if confirmar {
				w.Close()
			}
		}, w)
	})
	btnSalir.Importance = widget.MediumImportance
	btnSalir.Resize(fyne.NewSize(200, 50))

	// Contenedor del contenido
	contenido := container.NewVBox(
		container.NewCenter(titulo),
		container.NewCenter(subtitulo),
		container.NewCenter(linea),
		container.NewCenter(btnLogin),
		container.NewCenter(btnSalir),
	)

	// Layout final con fondo
	return container.NewMax(
		bg,
		container.NewCenter(contenido),
	)
}

func BuildLoginUI(w fyne.Window) fyne.CanvasObject {
	usuario := widget.NewEntry()
	pass := widget.NewPasswordEntry()

	// Acción común de login: valida campos, llama al controlador y navega al dashboard
	loginAction := func() {
		uText := strings.TrimSpace(usuario.Text)
		pText := pass.Text

		if uText == "" || pText == "" {
			dialog.ShowInformation("Campos incompletos", "Por favor ingresa usuario y contraseña.", w)
			return
		}

		user, err := controllers.VerificarCredenciales(uText, pText)
		if err != nil {
			dialog.ShowInformation("Error de autenticación", err.Error(), w)
			pass.SetText("")
			return
		}

		// Credenciales válidas: abrir dashboard con el usuario autenticado
		w.SetContent(BuildDashboardUI(w, user))
	}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: usuario},
			{Text: "Contraseña", Widget: pass},
		},
		OnSubmit: loginAction,
	}

	btnSubmit := widget.NewButtonWithIcon("Iniciar Sesión", theme.LoginIcon(), loginAction)
	btnSubmit.Importance = widget.HighImportance

	btnSalir := widget.NewButtonWithIcon("Salir", theme.CancelIcon(), func() {
		w.Close() // Cierra la ventana actual
	})

	// Contenedor del 'card' de login con tamaño proporcional a la ventana
	title := canvas.NewText("Iniciar Sesión", color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Determinar tamaño proporcional (fallback si canvas no está listo aún)
	canvasSize := w.Canvas().Size()
	if canvasSize.Width == 0 || canvasSize.Height == 0 {
		canvasSize = fyne.NewSize(600, 400)
	}
	cardW := int(canvasSize.Width * 0.45)
	cardH := int(canvasSize.Height * 0.5)

	// Construir un "card" manual usando un rectángulo oscuro con tamaño mínimo
	bgCard := canvas.NewRectangle(color.NRGBA{R: 30, G: 30, B: 35, A: 255})
	bgCard.SetMinSize(fyne.NewSize(float32(cardW), float32(cardH)))

	cardInner := container.NewVBox(
		container.NewCenter(title),
		container.NewPadded(form),
		container.NewHBox(btnSubmit, btnSalir),
	)

	cardWrap := container.NewMax(bgCard, container.NewPadded(cardInner))

	// Fondo oscuro para centrar mejor el card
	bg := canvas.NewRectangle(color.NRGBA{R: 20, G: 20, B: 25, A: 255})

	return container.NewMax(
		bg,
		container.NewCenter(cardWrap),
	)

}
