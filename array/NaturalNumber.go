package array

import "fmt"

func Natural(a,b int)  {
	sum:=0
	for i := a; i <=b; i++ {
		sum+=i
	}
	fmt.Println(sum)
}