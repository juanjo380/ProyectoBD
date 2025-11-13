package main

import (
	"fmt"
	"log"

	"ProyectoBD/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Conexión a PostgreSQL")

	statusLabel := widget.NewLabel("Verificando conexión...")

	// Llamamos correctamente a la función que devuelve (db, error)
	conn, err := db.Connect()

	if err != nil {
		log.Println("❌ Error:", err)
		statusLabel.SetText(fmt.Sprintf("❌ Error al conectar:\n%v", err))
	} else {
		defer conn.Close() // usamos conn.Close(), no db.CloseDB()
		statusLabel.SetText("✅ Conexión exitosa a PostgreSQL")
	}

	w.SetContent(container.NewVBox(
		widget.NewLabel("Estado de la conexión:"),
		statusLabel,
	))

	w.Resize(fyne.NewSize(400, 200))
	w.ShowAndRun()
}
