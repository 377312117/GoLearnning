package concurrent_learnning

import "fmt"

//
//  counter
//  @Description: out 代表着ch, int 发送给out,
//  @param out
//
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

//
//  squarer
//  @Description: in 代表着ch,表示接受型的channel，
//  @param out
//  @param in
//
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Code3() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
