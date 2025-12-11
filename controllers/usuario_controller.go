package controllers

import (
	"ProyectoBD/db"
	"ProyectoBD/models"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Obtener todos los usuarios
func GetAllUsuarios() ([]models.Usuario, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
        SELECT idUsuario, username, rol, nombreCompleto
        FROM Usuario
        ORDER BY idUsuario
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Usuario

	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(&u.IDUsuario, &u.Username, &u.Rol, &u.NombreCompleto)
		if err != nil {
			return nil, err
		}
		lista = append(lista, u)
	}

	return lista, nil
}

// Insertar un usuario
func InsertUsuario(u models.Usuario) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Encriptar contraseña antes de guardar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = conn.Exec(`
        INSERT INTO Usuario (username, password, rol, nombreCompleto)
        VALUES ($1, $2, $3, $4)
    `, u.Username, string(hashedPassword), u.Rol, u.NombreCompleto)

	if err != nil {
		fmt.Println("ERROR INSERTANDO USUARIO:", err)
	} else {
		fmt.Println("USUARIO INSERTADO OK")
	}

	return err
}

// Actualizar un usuario
func UpdateUsuario(u models.Usuario) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Si hay nueva contraseña, encriptarla
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = conn.Exec(`
            UPDATE Usuario
            SET username=$1, password=$2, rol=$3, nombreCompleto=$4
            WHERE idUsuario=$5
        `, u.Username, string(hashedPassword), u.Rol, u.NombreCompleto, u.IDUsuario)

		if err != nil {
			fmt.Println("ERROR ACTUALIZANDO USUARIO:", err)
		} else {
			fmt.Println("USUARIO ACTUALIZADO OK")
		}
	} else {
		// Actualizar sin cambiar contraseña
		_, err = conn.Exec(`
            UPDATE Usuario
            SET username=$1, rol=$2, nombreCompleto=$3
            WHERE idUsuario=$4
        `, u.Username, u.Rol, u.NombreCompleto, u.IDUsuario)

		if err != nil {
			fmt.Println("ERROR ACTUALIZANDO USUARIO:", err)
		} else {
			fmt.Println("USUARIO ACTUALIZADO OK")
		}
	}

	return err
}

// Eliminar un usuario
func DeleteUsuario(id int) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DELETE FROM Usuario WHERE idUsuario = $1", id)

	if err != nil {
		fmt.Println("ERROR ELIMINANDO USUARIO:", err)
	} else {
		fmt.Println("USUARIO ELIMINADO OK")
	}

	return err
}

// Obtener un usuario por ID
func GetUsuarioByID(id int) (*models.Usuario, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var u models.Usuario

	err = conn.QueryRow(`
        SELECT idUsuario, username, rol, nombreCompleto
        FROM Usuario
        WHERE idUsuario = $1
    `, id).Scan(&u.IDUsuario, &u.Username, &u.Rol, &u.NombreCompleto)

	if err != nil {
		fmt.Println("ERROR BUSCANDO USUARIO:", err)
		return nil, err
	}

	return &u, nil
}

// Obtener un usuario por username
func GetUsuarioByUsername(username string) (*models.Usuario, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var u models.Usuario
	var password string

	err = conn.QueryRow(`
        SELECT idUsuario, username, password, rol, nombreCompleto
        FROM Usuario
        WHERE username = $1
    `, username).Scan(&u.IDUsuario, &u.Username, &password, &u.Rol, &u.NombreCompleto)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	u.Password = password // Guardar contraseña encriptada
	return &u, nil
}

// Verificar credenciales de usuario

/*
func VerificarCredenciales(username, password string) (*models.Usuario, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var u models.Usuario
	var dbPassword string

	err = conn.QueryRow(`
        SELECT idUsuario, username, password, rol, nombreCompleto
        FROM Usuario
        WHERE username = $1`,
		username,
	).Scan(&u.IDUsuario, &u.Username, &dbPassword, &u.Rol, &u.NombreCompleto)

	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("contraseña incorrecta")
	}

	return &u, nil
}

*/

// Esta función DEBE existir en tu archivo
func VerificarCredenciales(username, password string) (*models.Usuario, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var u models.Usuario
	var dbPassword string

	err = conn.QueryRow(`
        SELECT idUsuario, username, password, rol, nombreCompleto
        FROM Usuario 
        WHERE username = $1`,
		username,
	).Scan(&u.IDUsuario, &u.Username, &dbPassword, &u.Rol, &u.NombreCompleto)

	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	// Verificar contraseña con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("contraseña incorrecta")
	}

	return &u, nil
}

// Cambiar contraseña de usuario
func CambiarPassword(id int, nuevaPassword string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nuevaPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = conn.Exec(`
        UPDATE Usuario
        SET password = $1
        WHERE idUsuario = $2
    `, string(hashedPassword), id)

	if err != nil {
		fmt.Println("ERROR CAMBIANDO PASSWORD:", err)
	} else {
		fmt.Println("PASSWORD CAMBIADA OK")
	}

	return err
}
