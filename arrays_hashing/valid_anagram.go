package array_shashing

func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	//Guardar el primer string en un `mapa[caracter] cantidad`
	characters_amount := make(map[rune]int)
	for _, c := range s {
		characters_amount[c]++
	}

	//Recorrer el segundo string disminiyendo las cantidades en el mapa,
	for _, c := range t {
		_, exists := characters_amount[c]

		//validaciones diferenciales
		if !exists || characters_amount[c] == 0 {
			return false
		} else {
			characters_amount[c]--
		}
	}

	return true
}
