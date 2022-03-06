package interface_learnning

import "fmt"

type Service interface {
	Start()     // 开启服务
	Log(string) // 打印日志
}

type Logger struct {
}

func (g *Logger) Log(s string) {
	fmt.Println("println log:", s)
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {
	fmt.Println("Game Start...")
}

func Code2() {
	var a Service = new(GameService)
	a.Start()
	a.Log("hello")
}
