package libros

import (
	"errors"
	"fmt"
	"strings"
)

type Libro struct {
	id            int
	titulo        string
	autor         string
	categoria     string
	anio          int
	codigo        string
	disponible    bool
}

func NewLibro(id int, titulo, autor, categoria string, anio int, codigo string) (*Libro, error) {
	if id <= 0 {
		return nil, errors.New("el ID del libro debe ser mayor que cero")
	}

	if strings.TrimSpace(titulo) == "" {
		return nil, errors.New("el título no puede estar vacío")
	}

	if strings.TrimSpace(autor) == "" {
		return nil, errors.New("el autor no puede estar vacío")
	}

	if strings.TrimSpace(categoria) == "" {
		return nil, errors.New("la categoría no puede estar vacía")
	}

	if anio <= 0 {
		return nil, errors.New("el año debe ser válido")
	}

	if strings.TrimSpace(codigo) == "" {
		return nil, errors.New("el código no puede estar vacío")
	}

	return &Libro{
		id:         id,
		titulo:     titulo,
		autor:      autor,
		categoria:  categoria,
		anio:       anio,
		codigo:     codigo,
		disponible: true,
	}, nil
}

func (l Libro) GetID() int {
	return l.id
}

func (l Libro) GetTitulo() string {
	return l.titulo
}

func (l Libro) GetAutor() string {
	return l.autor
}

func (l Libro) GetCategoria() string {
	return l.categoria
}

func (l Libro) GetAnio() int {
	return l.anio
}

func (l Libro) GetCodigo() string {
	return l.codigo
}

func (l Libro) EstaDisponible() bool {
	return l.disponible
}

func (l *Libro) SetTitulo(titulo string) error {
	if strings.TrimSpace(titulo) == "" {
		return errors.New("el título no puede estar vacío")
	}

	l.titulo = titulo
	return nil
}

func (l *Libro) SetCategoria(categoria string) error {
	if strings.TrimSpace(categoria) == "" {
		return errors.New("la categoría no puede estar vacía")
	}

	l.categoria = categoria
	return nil
}

func (l *Libro) Prestar() error {
	if !l.disponible {
		return errors.New("el libro no está disponible")
	}

	l.disponible = false
	return nil
}

func (l *Libro) Devolver() {
	l.disponible = true
}

func (l Libro) MostrarInfo() string {
	return fmt.Sprintf("Libro: %s - Autor: %s - Categoría: %s",
		l.titulo, l.autor, l.categoria)
}
