package main

import (
	"fmt"

	"github.com/alevale2004/gestion-libros-electronicos-go/autores"
	"github.com/alevale2004/gestion-libros-electronicos-go/interfaces"
	"github.com/alevale2004/gestion-libros-electronicos-go/libros"
	"github.com/alevale2004/gestion-libros-electronicos-go/usuarios/usuarios"
)

// Esta función puede recibir cualquier elemento que implemente la interfaz Mostrable.
func mostrarElemento(elemento interfaces.Mostrable) {
	fmt.Println(elemento.MostrarInfo())
}

func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos")
	fmt.Println("-----------------------------------------")

	usuario, err := usuarios.NewUsuario(
		1,
		"Alexandra",
		"alexandra@email.com",
		"1723456789",
		"Estudiante",
	)

	if err != nil {
		fmt.Println("Error al crear usuario:", err)
		return
	}

	libro, err := libros.NewLibro(
		1,
		"El Principito",
		"Antoine de Saint-Exupéry",
		"Literatura",
		1943,
		"LIB001",
	)

	if err != nil {
		fmt.Println("Error al crear libro:", err)
		return
	}

	autor, err := autores.NewAutor(
		1,
		"Antoine de Saint-Exupéry",
		"Francés",
		"Autor y aviador francés",
	)

	if err != nil {
		fmt.Println("Error al crear autor:", err)
		return
	}

	fmt.Println("\nInformación registrada:")

	mostrarElemento(usuario)
	mostrarElemento(libro)
	mostrarElemento(autor)
}
