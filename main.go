package main

import (
	"fmt"
	"github.com/m/v1/zzx/gocode/interface_test"
)

func test(an interfacetest.Animal) {
	fmt.Printf("animals:%s, %s", an.Name(), an.Fly())
}

func main() {
	fish1 := interfacetest.Fish{}
	test(fish1)
}
