package main

import (
	"ProyectoBD/controllers"
	"ProyectoBD/db"
	"ProyectoBD/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Inicializar la base de datos
	db.Init()

	// Asegurar superadmin por defecto
	_ = controllers.EnsureDefaultSuperAdmin()

	// Crear la app
	a := app.New()

	// Crear la ventana principal
	w := a.NewWindow("Sistema de Confecciones – Login")

	w.SetContent(ui.BuildWelcomeUI(w))
	// Mostrar pantalla de login

	// Tamaño inicial
	w.Resize(fyne.NewSize(600, 400))

	// Ejecutar
	w.ShowAndRun()
}
