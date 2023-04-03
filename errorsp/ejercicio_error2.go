package errorsp

type NuevoError struct {
	Mensaje string
}

func (n *NuevoError) Error() string {
	return n.Mensaje
}

var ErrImpSala = &NuevoError{Mensaje: "Error: el salario es menor a 10.000"}

func ValidarSalario(salario int) error {
	if salario <= 10000 {
		return ErrImpSala
	}
	return nil
}
