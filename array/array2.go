package array

func Pifagor(a, b, c int) bool {
	if a >= b+c || b >= c+a || c >= a+b {
		return false
	} else {
		return a*a+b*b == c*c
	}

}
