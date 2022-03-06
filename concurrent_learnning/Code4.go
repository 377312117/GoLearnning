package concurrent_learnning

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func GoTest1(ctx context.Context, wg *sync.WaitGroup) {
loop:
	for {
		time.Sleep(2 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("select Done!")
			break loop
		default:
			fmt.Println("goroutine default running...")
		}
	}
	wg.Done()
	fmt.Println("wg is Done!")
}

func Code4() {
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go GoTest1(ctx, &wg)
mainLoop:
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("for main loop is running...")
		select {
		case sig := <-exitChan:
			fmt.Println("接受到来自系统的信号：", sig)
			//os.Exit(1) //如果ctrl+c 关不掉程序，使用os.Exit强行关掉
			break mainLoop
		}
	}
	cancel()
	wg.Wait()
	fmt.Println("main goroutine is Done!")
}
