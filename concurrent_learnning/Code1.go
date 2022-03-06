package concurrent_learnning

import (
	"fmt"
	"sync"
)

func Code1() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}
