package array

func White(num []int) []int {
	Black := []int{}
	for _, v := range num {
		if v >= 0 && v <= 100 {
			Black = append(Black, v)
		} else if v > 100 {
			Black = append(Black, v+v)
		}
	}
	return Black
}
