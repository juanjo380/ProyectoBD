package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	IDUsuario      int    `json:"idUsuario"`
	Username       string `json:"username"`
	Password       string `json:"-"`
	Rol            string `json:"rol"` // "administrador" o "vendedor"
	NombreCompleto string `json:"nombreCompleto"`
}

// HashPassword encripta la contraseña
func (u *Usuario) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

// CheckPassword verifica la contraseña
func (u *Usuario) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// EsAdministrador verifica si es administrador
func (u *Usuario) EsAdministrador() bool {
	return u.Rol == "administrador"
}

// EsVendedor verifica si es vendedor
func (u *Usuario) EsVendedor() bool {
	return u.Rol == "vendedor"
}
