package ui

import (
	"ProyectoBD/controllers"
	"image/color"

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
		w.Resize(fyne.NewSize(500, 400))
	})
	btnLogin.Importance = widget.HighImportance
	btnLogin.Resize(fyne.NewSize(200, 50))

	// Contenedor del contenido
	contenido := container.NewVBox(
		container.NewCenter(titulo),
		container.NewCenter(subtitulo),
		container.NewCenter(linea),
		container.NewCenter(btnLogin),
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

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: usuario},
			{Text: "Contraseña", Widget: pass},
		},
		OnSubmit: func() {
			// Verificar credenciales
			_, err := controllers.VerificarCredenciales(usuario.Text, pass.Text)
			if err != nil {
				dialog.ShowInformation("Error", "Usuario o contraseña incorrectos", w)
				return
			}

			// Si las credenciales son correctas, pasar al dashboard
			w.SetContent(BuildDashboardUI(w))
		},
	}

	return container.NewVBox(
		widget.NewLabelWithStyle("Iniciar Sesión",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		form,
	)
}

/*

// Pantalla de inicio de sesión
func BuildLoginUI(w fyne.Window) fyne.CanvasObject {

	usuario := widget.NewEntry()
	pass := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: usuario},
			{Text: "Contraseña", Widget: pass},
		},
		OnSubmit: func() {

			// Aquí defines el login, hardcodeado
			if usuario.Text == "admin" && pass.Text == "1234" {

				// Si todo está ok → pasar al dashboard
				w.SetContent(BuildDashboardUI(w))
				return
			}

			dialog.ShowInformation("Error", "Usuario o contraseña incorrectos", w)
		},
	}

	return container.NewVBox(
		widget.NewLabelWithStyle("Iniciar Sesión",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		form,
	)
}*/
