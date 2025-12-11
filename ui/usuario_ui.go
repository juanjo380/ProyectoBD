package ui

import (
	"ProyectoBD/controllers"
	"ProyectoBD/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildUsuarioUI(w fyne.Window) fyne.CanvasObject {
	// Obtener usuarios
	usuarios, _ := controllers.GetAllUsuarios()

	// Crear tabla
	data := [][]string{{"ID", "Usuario", "Rol", "Nombre Completo"}}
	for _, u := range usuarios {
		data = append(data, []string{
			strconv.Itoa(u.IDUsuario),
			u.Username,
			u.Rol,
			u.NombreCompleto,
		})
	}

	table := widget.NewTable(
		func() (int, int) { return len(data), 4 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(cell widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cell.Row][cell.Col])
		},
	)

	// Botones
	btnAdd := widget.NewButton("Nuevo Usuario", func() {
		w.SetContent(BuildCrearUsuarioUI(w))
	})

	btnEdit := widget.NewButton("Editar Usuario", func() {
		openUsuarioEditDialog(w)
	})

	btnDelete := widget.NewButton("Eliminar Usuario", func() {
		openUsuarioDeleteDialog(w)
	})

	//boton salir
	btnSalir := widget.NewButton("Salir", func() {
		w.SetContent(BuildDashboardUI(w, nil))
	})

	// Solo admin puede crear/editar usuarios
	// Aquí podrías agregar lógica para verificar rol

	menu := container.NewVBox(btnAdd, btnEdit, btnDelete, btnSalir)

	return container.NewBorder(menu, nil, nil, nil, table)
}

func BuildCrearUsuarioUI(w fyne.Window) fyne.CanvasObject {
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	nombreCompleto := widget.NewEntry()

	rolSelect := widget.NewSelect([]string{"administrador", "vendedor"}, nil)
	rolSelect.SetSelected("vendedor")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: username},
			{Text: "Contraseña", Widget: password},
			{Text: "Nombre Completo", Widget: nombreCompleto},
			{Text: "Rol", Widget: rolSelect},
		},
		OnSubmit: func() {
			nuevoUsuario := models.Usuario{
				Username:       username.Text,
				Password:       password.Text,
				Rol:            rolSelect.Selected,
				NombreCompleto: nombreCompleto.Text,
			}

			err := controllers.InsertUsuario(nuevoUsuario)
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				w.SetContent(BuildUsuarioUI(w))
			}
		},
		OnCancel: func() {
			w.SetContent(BuildUsuarioUI(w))
		},
	}

	return container.NewVBox(widget.NewLabel("Registrar Nuevo Usuario"), form)
}

func openUsuarioEditDialog(w fyne.Window) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID del usuario")

	dialog.ShowForm("Editar Usuario",
		"Buscar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID del Usuario", Widget: idEntry},
		},
		func(ok bool) {
			if !ok {
				return
			}

			id, err := strconv.Atoi(idEntry.Text)
			if err != nil {
				dialog.ShowInformation("Error", "ID inválido", w)
				return
			}

			usuario, err := controllers.GetUsuarioByID(id)
			if err != nil {
				dialog.ShowInformation("Error", "Usuario no encontrado", w)
				return
			}

			w.SetContent(BuildEditarUsuarioUI(w, usuario))
		},
		w,
	)
}

func BuildEditarUsuarioUI(w fyne.Window, u *models.Usuario) fyne.CanvasObject {
	username := widget.NewEntry()
	username.SetText(u.Username)

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Dejar vacío para no cambiar")

	nombreCompleto := widget.NewEntry()
	nombreCompleto.SetText(u.NombreCompleto)

	rolSelect := widget.NewSelect([]string{"administrador", "vendedor"}, nil)
	rolSelect.SetSelected(u.Rol)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Usuario", Widget: username},
			{Text: "Nueva Contraseña (opcional)", Widget: password},
			{Text: "Nombre Completo", Widget: nombreCompleto},
			{Text: "Rol", Widget: rolSelect},
		},
		OnSubmit: func() {
			u.Username = username.Text
			u.NombreCompleto = nombreCompleto.Text
			u.Rol = rolSelect.Selected

			// Solo actualizar password si se ingresó una nueva
			if password.Text != "" {
				u.Password = password.Text
			}

			err := controllers.UpdateUsuario(*u)
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				w.SetContent(BuildUsuarioUI(w))
			}
		},
		OnCancel: func() {
			w.SetContent(BuildUsuarioUI(w))
		},
	}

	return container.NewVBox(
		widget.NewLabel("Editando Usuario: "+u.Username),
		form,
	)
}

func openUsuarioDeleteDialog(w fyne.Window) {
	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("ID del usuario")

	dialog.ShowForm("Eliminar Usuario",
		"Eliminar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "ID del Usuario", Widget: idEntry},
		},
		func(ok bool) {
			if !ok {
				return
			}

			id, err := strconv.Atoi(idEntry.Text)
			if err != nil {
				dialog.ShowInformation("Error", "ID inválido", w)
				return
			}

			controllers.DeleteUsuario(id)
			w.SetContent(BuildUsuarioUI(w))
		},
		w)
}
