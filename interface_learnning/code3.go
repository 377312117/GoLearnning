package interface_learnning

import "fmt"

func Code5() int {
	x := 0
	defer func() {
		// x := x
		fmt.Println("x:", x)
	}()
	x = 1
	return x
}
