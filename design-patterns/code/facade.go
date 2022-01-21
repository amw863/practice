package code

import "fmt"

// 该角色对外提供服务
// 也是一个客服端
// 外部可以直接调用子系统，子系统并非一个方法

type Opener interface {
	Open()
}

type Tv struct {
}

func (receiver Tv) Open() {
	fmt.Println("open tv")
}

type Light struct {
}

func (l Light) Open() {
	fmt.Println("open light")
}

type Heater struct {
}

func (h Heater) Open() {
	fmt.Println("open heater")
}

// Facade 门面模式
//优势：对客户屏蔽了子系统组件，客户端代码边的简单
//劣势：违反了开闭原则
type Facade struct {
	t Tv
	l Light
	h Heater
}

func (f Facade) Open() {
	f.t.Open()
	f.l.Open()
	f.h.Open()
}
