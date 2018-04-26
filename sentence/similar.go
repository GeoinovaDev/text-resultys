package sentence

import "strings"

func similar(sentence1 string, sentence2 string) float32 {
	sentence1 = clean(sentence1)
	sentence2 = clean(sentence2)

	if len(sentence1) == 0 || len(sentence2) == 0 {
		return 0
	}

	arr1 := strings.Split(sentence1, " ")
	arr2 := strings.Split(sentence2, " ")
	swap := arr2

	if len(arr1) > len(arr2) {
		arr2 = arr1
		arr1 = swap
	}

	counter := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				counter++
				break
			}
		}
	}

	return float32(counter / len(arr1))
}
