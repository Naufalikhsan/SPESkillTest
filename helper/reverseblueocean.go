package helper

func ReverseBlueOcean(arrayInt1 []int, arrayInt2 []int) []int {
	var result []int
	for _, val := range arrayInt1 {
		if val != arrayInt2[0] {
			result = append(result, val)
		}
	}

	return result
}
