package helper

func ParityOutlier(arrayInt []int) (num int, isoutlier bool) {

	for i := range arrayInt {
		number := arrayInt[i]
		if number%2 == 0 {
			num = number
		}
	}

	if num == 0 {
		isoutlier = false
	}

	return num, isoutlier
}
