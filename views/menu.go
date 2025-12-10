package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MenuView(a fyne.App, w fyne.Window) fyne.CanvasObject {

	label := widget.NewLabel("Bienvenido al sistema de confecciones")

	salirBtn := widget.NewButton("Cerrar aplicación", func() {
		a.Quit()
	})

	return container.NewVBox(
		label,
		salirBtn,
	)
}
