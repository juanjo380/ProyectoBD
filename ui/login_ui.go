package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

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
}

