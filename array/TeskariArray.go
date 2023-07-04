package array

func TeskarArray(str []int) []int {
	left := 0
	right := len(str) - 1
	for right > left {
		str[right], str[left] = str[left], str[right]
		left++
		right--
	}
	return str
}
