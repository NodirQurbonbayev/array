package array

import "math"

func SmallestMultipleOf(array []int, number1, number2 int) int {
	smallest := math.MaxInt32
	for _, num := range array {
	  if num%number1 == 0 && num%number2 != 0 {
		if num < smallest {
		  smallest = num
		}
	  }
	}
	return smallest
  }
  