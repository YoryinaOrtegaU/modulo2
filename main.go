package main

import l "pratica/panic_defer/defer1"

func main() {
	/* 	var v int = 19
	   	fmt.Println("la direccion de memoria de v es: ", &v)
	   	c := &v
	   	fmt.Println("el valor de c: ", c)
	   	fmt.Println("la direccion de memoria de c es: ", *c, v)
	   	*c = 20
	   	fmt.Println("valores c y v: ", *c, v)

	   	fmt.Println(f.CalcularImpuesto(160000))

	   	fmt.Println(f.CalcularPromedio(2, 3, 5, 5, 5, 4))

	   	fmt.Println(f.CalcularSalario(50, "C"))

	   	calculo := f.CalcularEstadisticas("promedio")
	   	resultado := calculo(1, 2, 3, 4, 5)

	   	fmt.Println(resultado)

	   	animal := f.Animal("gato")
	   	fmt.Println(animal(5))

	   	//Estructuras-structs

	   	product := e.Product{Id: 2,
	   		Name:        "leche",
	   		Price:       2500,
	   		Description: "leche deslactosada",
	   		Category:    "lacteos"}
	   	product2 := e.Product{Id: 3,
	   		Name:        "yogurt Mora",
	   		Price:       2500,
	   		Description: "yogurt sabor a mora",
	   		Category:    "lacteos"}

	   	var products []e.Product
	   	products = append(products, product, product2)
	   	product.Save(products)
	   	product.GetAll(products)
	   	fmt.Println(e.GetById(products, 3))

	   	persona := e.Person{Id: 01, Name: "Ricardo", DateOfBirth: "15 octubre"}
	   	empleado := e.Employee{Id: 02, Person: persona, Position: "Secretario"}

	   	empleado.PrintEmployee() */

	//Interfaces

	/* 	alumno := a.Alumno{Nombre: "Julio", Apellido: "Pello", Dni: "123", Fecha: "25 agosto"}
	   	alumno.Detalle()
	   	alumno.Detalle2()

	   	product := a.FactoryP("Grande", 20000)
	   	precioMediano := product.PrecioTotal()
	   	fmt.Println(precioMediano) */

	//errors
	/* 	salary := 10000
	   	fmt.Println(err.CalcularImpuesto(salary))

	   	resultado := err.ValidarSalario(salary)
	   	conicidencia := errors.Is(resultado, err.ErrImpSala)
	   	fmt.Println("error es:", conicidencia)

	   	resultado3 := err.CalcularImpuesto2(salary)
	   	conicidencia2 := errors.Is(resultado3, err.Error1)
	   	fmt.Println("error nuevo es:", conicidencia2)

	   	resultado4 := err.CalcularImpuesto4(salary)
	   	fmt.Println(resultado4)

	   	salario, posibleError := err.CalcularSalario5(100, 30000)

	   	if posibleError != nil {
	   		fmt.Println(posibleError)
	   	} else {
	   		fmt.Println(salario)
	   	} */
	l.LeerArchivo("archivo")
}
