package usuarios

import (
	"errors"
	"strings"
)

type Usuario struct {
	id             int
	nombre         string
	correo         string
	identificacion string
	tipo           string
}

func NewUsuario(id int, nombre, correo, identificacion, tipo string) (*Usuario, error) {
	if id <= 0 {
		return nil, errors.New("el ID del usuario debe ser mayor que cero")
	}

	if strings.TrimSpace(nombre) == "" {
		return nil, errors.New("el nombre del usuario no puede estar vacío")
	}

	if strings.TrimSpace(correo) == "" {
		return nil, errors.New("el correo del usuario no puede estar vacío")
	}

	if strings.TrimSpace(identificacion) == "" {
		return nil, errors.New("la identificación del usuario no puede estar vacía")
	}

	if strings.TrimSpace(tipo) == "" {
		return nil, errors.New("el tipo de usuario no puede estar vacío")
	}

	return &Usuario{
		id:             id,
		nombre:         nombre,
		correo:         correo,
		identificacion: identificacion,
		tipo:           tipo,
	}, nil
}

func (u Usuario) GetID() int {
	return u.id
}

func (u Usuario) GetNombre() string {
	return u.nombre
}

func (u Usuario) GetCorreo() string {
	return u.correo
}

func (u Usuario) GetIdentificacion() string {
	return u.identificacion
}

func (u Usuario) GetTipo() string {
	return u.tipo
}


func (u *Usuario) SetNombre(nombre string) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	u.nombre = nombre
	return nil
}

func (u *Usuario) SetCorreo(correo string) error {
	if strings.TrimSpace(correo) == "" {
		return errors.New("el correo no puede estar vacío")
	}

	u.correo = correo
	return nil
}

func (u *Usuario) SetTipo(tipo string) error {
	if strings.TrimSpace(tipo) == "" {
		return errors.New("el tipo de usuario no puede estar vacío")
	}

	u.tipo = tipo
	return nil
}
