package prestamos

import (
	"errors"
	"time"
)

type Prestamo struct {
	id              int
	idUsuario       int
	idLibro         int
	fechaPrestamo   time.Time
	fechaDevolucion time.Time
	estado          string
}

func NewPrestamo(id, idUsuario, idLibro int) (*Prestamo, error) {
	if id <= 0 {
		return nil, errors.New("el ID del préstamo debe ser mayor que cero")
	}

	if idUsuario <= 0 {
		return nil, errors.New("el ID del usuario no es válido")
	}

	if idLibro <= 0 {
		return nil, errors.New("el ID del libro no es válido")
	}

	return &Prestamo{
		id:            id,
		idUsuario:     idUsuario,
		idLibro:       idLibro,
		fechaPrestamo: time.Now(),
		estado:        "Activo",
	}, nil
}

func (p Prestamo) GetID() int {
	return p.id
}

func (p Prestamo) GetIDUsuario() int {
	return p.idUsuario
}

func (p Prestamo) GetIDLibro() int {
	return p.idLibro
}

func (p Prestamo) GetFechaPrestamo() time.Time {
	return p.fechaPrestamo
}

func (p Prestamo) GetFechaDevolucion() time.Time {
	return p.fechaDevolucion
}

func (p Prestamo) GetEstado() string {
	return p.estado
}

func (p *Prestamo) Devolver() error {
	if p.estado == "Devuelto" {
		return errors.New("el préstamo ya fue devuelto")
	}

	p.fechaDevolucion = time.Now()
	p.estado = "Devuelto"

	return nil
}
