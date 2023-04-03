package errorsp

import "errors"

// Ejercicio 1
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func CalcularImpuesto(salary int) (string, error) {
	if salary < 150000 {
		return "", &MyError{msg: "el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return "Debe pagar impuesto", nil
	}
}

//Ejercicio 2
/* En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error:
el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
La validación debe ser hecha con la función “Is()” dentro del “main”. */

type ErrorSalario struct {
	Mensaje string
}

func (e *ErrorSalario) Error() string {
	return e.Mensaje
}

func aplicaImpuesto(salario int) (string, error) {
	if salario < 150000 {
		return "", ErrTaxSalary
	}
	return "Debe pagar impuesto", nil
}

var (
	ErrTaxSalary      = &ErrorSalario{Mensaje: "Error: el salario ingresado no alcanza el mínimo imponible"}
	ErrSalaryQuantity = &ErrorSalario{Mensaje: "Error: el salario es menor a 10.000"}
)

func validarMontoSalario(salario int) error {
	if salario <= 10000 {
		return ErrSalaryQuantity
	}
	return nil
}

func main2() {
	salario := 100
	aplica, errImpuesto := aplicaImpuesto(salario)
	errMontoSalario := validarMontoSalario(salario)

	if errImpuesto != nil {
		println(errImpuesto.Error())
	}

	if errors.Is(errMontoSalario, ErrSalaryQuantity) {
		println(errMontoSalario.Error())
	}

	println(aplica)
}
