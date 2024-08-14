package array_shashing

import "sort"

type kv struct {
	Key   int
	Value int
}

func TopKFrequent(nums []int, k int) []int {
	frequentMap := make(map[int]int)

	for _, value := range nums {
		frequentMap[value]++
	}

	// Slice para almacenar los pares clave-valor
	var kvs []kv
	for k, v := range frequentMap {
		kvs = append(kvs, kv{k, v})
	}

	// Ordenar el slice por los valores
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	// Crear un slice para almacenar las claves ordenadas
	var keys []int
	for _, kv := range kvs {
		keys = append(keys, kv.Key)
	}

	return keys[:k]
}
