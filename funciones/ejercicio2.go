package funciones

func CalcularPromedio(calificaciones ...uint) uint {
	var promedio uint
	var sumaNotas uint

	for _, calificacion := range calificaciones {

		sumaNotas += calificacion
	}
	promedio = sumaNotas / uint(len(calificaciones))
	return promedio
}
