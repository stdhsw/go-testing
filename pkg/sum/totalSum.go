package sum

func InefficientSum(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}

func EfficientSum(num int) int {
	return (num * (num + 1)) / 2
}
