package soma

//Soma - +, +, + ...
func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// func Soma(numeros [5]int) int {
// 	soma := 0
// 	for i := 0; i < len(numeros); i++ {
// 		soma += numeros[i]
// 	}
// 	return soma
// }

//Tudo - de slice em outro slice
func Tudo(slicesTotal ...[]int) []int {
	var somas []int
	for _, numeros := range slicesTotal {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

//TodoResto - menos o inicial
func TodoResto(totalSlice ...[]int) []int {
	var somas []int
	for _, numeros := range totalSlice {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	return somas
}
