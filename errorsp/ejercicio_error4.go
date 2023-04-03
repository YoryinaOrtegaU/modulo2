package errorsp

import (
	"errors"
	"fmt"
)

var Error4 = errors.New("el salario es menor a 10.000")

func CalcularImpuesto4(salary int) error {
	if salary <= 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}
