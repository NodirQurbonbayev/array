package array

func YearsToGrow(x float64, p float64, y float64) int {
	years := 0
	currentAmount := x
	for currentAmount < y {
	  currentAmount = currentAmount * (1 + p/100)
	  years++
	}
	return years
  }
  