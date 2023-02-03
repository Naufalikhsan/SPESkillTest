package helper

func IntToSlice(number int) []int {
	slice := []int{}

	for number > 0 {
		slice = append(slice, number%10)
		number /= 10
	}

	result := []int{}
	for i := range slice {
		result = append(result, slice[len(slice)-1-i])
	}

	return result
}
