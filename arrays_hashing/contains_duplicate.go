package array_shashing

func ContainsDuplicate(nums []int) bool {

	valuesAmount := make(map[int]bool)
	result := false

	//Ir guardando el array de numeros en la key de un mapa
	for _, v := range nums {

		//key ya existe es que esta duplicada en el array
		if _, found := valuesAmount[v]; found {
			result = true
			break
		} else {
			valuesAmount[v] = true
		}
	}

	return result
}
