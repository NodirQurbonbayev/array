package array

func List(a, b int) []int {
	sum := []int{}
	for i := a; i <= b; i++ {
		sum = append(sum, i*i)
	}
	return sum
}
