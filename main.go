package main

import (
    "ProyectoBD/ui"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
)

func main() {

    // Crear la app
    a := app.New()

    // Crear la ventana principal
    w := a.NewWindow("Sistema de Confecciones – Login")

    // Mostramos primero la pantalla de login
    w.SetContent(ui.BuildLoginUI(w))

    // Tamaño inicial
    w.Resize(fyne.NewSize(500, 300))

    // Ejecutar
    w.ShowAndRun()
}


