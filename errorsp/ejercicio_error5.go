package errorsp

type ErroSalario struct {
	Mensaje string
}

func (e *ErroSalario) Error() string {
	return e.Mensaje
}

func CalcularSalario5(horasTrabajadas float64, valorHora float64) (float64, error) {

	if horasTrabajadas < 80 {
		var EImpuesto = &ErroSalario{Mensaje: "Error: el trabajador no puede haber trabajado menos de 80 hs mensuales"}
		return 0.0, EImpuesto
	}

	salarioCalculado := horasTrabajadas * valorHora

	if salarioCalculado >= 150000 {
		impuesto := salarioCalculado * 0.10
		salarioCalculado = salarioCalculado - impuesto
		return salarioCalculado, nil
	}
	return salarioCalculado, nil
}
