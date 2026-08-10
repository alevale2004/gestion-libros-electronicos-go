package autores

import (
	"errors"
	"fmt"
	"strings"
)

type Autor struct {
	id           int
	nombre       string
	nacionalidad string
	biografia    string
}

func NewAutor(id int, nombre, nacionalidad, biografia string) (*Autor, error) {
	if id <= 0 {
		return nil, errors.New("el ID del autor debe ser mayor que cero")
	}

	if strings.TrimSpace(nombre) == "" {
		return nil, errors.New("el nombre del autor no puede estar vacío")
	}

	return &Autor{
		id:           id,
		nombre:       nombre,
		nacionalidad: nacionalidad,
		biografia:    biografia,
	}, nil
}

func (a Autor) GetID() int {
	return a.id
}

func (a Autor) GetNombre() string {
	return a.nombre
}

func (a Autor) GetNacionalidad() string {
	return a.nacionalidad
}

func (a Autor) GetBiografia() string {
	return a.biografia
}

func (a *Autor) SetNombre(nombre string) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	a.nombre = nombre
	return nil
}

func (a *Autor) SetNacionalidad(nacionalidad string) {
	a.nacionalidad = nacionalidad
}

func (a *Autor) SetBiografia(biografia string) {
	a.biografia = biografia
}

func (a Autor) MostrarInfo() string {
	return fmt.Sprintf("Autor: %s - Nacionalidad: %s",
		a.nombre, a.nacionalidad)
}
