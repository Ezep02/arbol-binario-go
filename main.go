package main

import "fmt"

type Nodo struct {
	Clave     int
	Derecho   *Nodo
	Izquierdo *Nodo
}

// insertar un nuevo valor
func (n *Nodo) Insert(valor int) {

	if valor < n.Clave {

		if n.Izquierdo == nil {
			n.Izquierdo = &Nodo{Clave: valor}
		} else {
			n.Izquierdo.Insert(valor)
		}
	} else {
		if n.Derecho == nil {
			n.Derecho = &Nodo{Clave: valor}
		} else {
			n.Derecho.Insert(valor)
		}
	}
}

// buscar un valor en el arbol
func (n *Nodo) Buscar(valor int) bool {

	if n == nil {
		return false
	}

	if valor < n.Clave {
		return n.Izquierdo.Buscar(valor)
	}
	if valor > n.Clave {
		return n.Derecho.Buscar(valor)
	}

	return true
}

func main() {

	// Crear la raiz del árbol binario
	root := &Nodo{Clave: 10}

	// Insertar valores en el árbol binario
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(17)

	// Buscar un valor en el arbol
	fmt.Println(root.Buscar(7))  // -> true
	fmt.Println(root.Buscar(20)) // -> false
}
