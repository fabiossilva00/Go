package iteracao

//Repetir - Repete 5x
func Repetir(caracter string, quantia int) string {
	var repeticoes string
	for i := 0; i < quantia; i++ {
		repeticoes += caracter
	}

	return repeticoes
}
