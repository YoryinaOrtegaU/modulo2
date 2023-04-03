package errorsp

import "errors"

var Error1 = errors.New("el salario es menor a 10.000")

func CalcularImpuesto2(salary int) error {
	if salary <= 10000 {
		return Error1
	}
	return nil
}
