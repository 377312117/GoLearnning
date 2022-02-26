package interfacetest

type Animal interface {
	Fly() string
	Swim() string
	Run() string
	Name() string
} 

type Dog struct {
}

func (d Dog) Run() string {
	return "can run"
}

func (d Dog) Swim() string {
	return "can swim"
}

func (d Dog) Fly() string {
	return "can fly"
}

func (d Dog) Name() string {
	return "dog"
}

type Fish struct {
}

func (f Fish) Run() string {
	return "can run"
}

func (f Fish) Swim() string {
	return "can swim"
}

func (f Fish) Fly() string {
	return "can fly"
}

func (f Fish) Name() string {
	return "fish"
}





