package categorias

import (
	"errors"
	"strings"
)

type Categoria struct {
	id          int
	nombre      string
	descripcion string
}

func NewCategoria(id int, nombre, descripcion string) (*Categoria, error) {
	if id <= 0 {
		return nil, errors.New("el ID de la categoría debe ser mayor que cero")
	}

	if strings.TrimSpace(nombre) == "" {
		return nil, errors.New("el nombre de la categoría no puede estar vacío")
	}

	return &Categoria{
		id:          id,
		nombre:      nombre,
		descripcion: descripcion,
	}, nil
}

func (c Categoria) GetID() int {
	return c.id
}

func (c Categoria) GetNombre() string {
	return c.nombre
}

func (c Categoria) GetDescripcion() string {
	return c.descripcion
}

func (c *Categoria) SetNombre(nombre string) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	c.nombre = nombre
	return nil
}

func (c *Categoria) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}
