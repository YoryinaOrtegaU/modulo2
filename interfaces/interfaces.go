package interfaces

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (a *Alumno) Detalle() {
	fmt.Printf("%+v\n", a)
}

func (a *Alumno) Detalle2() {
	fmt.Printf("Nombre: %s\n", a.Nombre)
	fmt.Printf("Apellido: %s\n", a.Apellido)
	fmt.Printf("Dni: %s\n", a.Dni)
	fmt.Printf("Fecha: %s\n", a.Fecha)
}

type Product interface {
	PrecioTotal() float64
}

type Pequeno struct {
	Precio float64
}

func (p Pequeno) PrecioTotal() float64 {
	var precioTotal float64 = p.Precio
	return precioTotal
}

type Mediano struct {
	Precio float64
}

func (m Mediano) PrecioTotal() float64 {
	precioTotal := m.Precio + m.Precio*0.03
	return precioTotal
}

type Grande struct {
	Precio float64
}

func (g Grande) PrecioTotal() float64 {
	precioTotal := g.Precio + g.Precio*0.06 + 2500
	return precioTotal
}

func FactoryP(tipoProducto string, precio float64) Product {
	switch tipoProducto {
	case "Pequeno":
		return Pequeno{Precio: precio}
	case "Mediano":
		return Mediano{Precio: precio}
	default:
		return Grande{Precio: precio}
	}
}
