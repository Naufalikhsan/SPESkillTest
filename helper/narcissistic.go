package helper

import "math"

func Narcissistic(val int) bool {
	digit := CountDigit(val)
	numSlice := IntToSlice(val)
	var sumPower int

	for i := range numSlice {
		number := numSlice[i]
		power := math.Pow(float64(number), float64(digit))
		//fmt.Println(number)
		//fmt.Println(x)
		sumPower += int(power)
	}

	if sumPower == val {
		return true
	} else {
		return false
	}
}

func CountDigit(number int) int {
	var count int
	if number > 0 {
		for number > 0 {
			number = number / 10
			count = count + 1
		}
	}

	return count
}
