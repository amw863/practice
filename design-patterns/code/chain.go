package code

import "fmt"

// 将请求和发送解耦，多个对象有激活处理当前请求
// 复用扩展在不改变
// 降低耦合度/简化对象互相链接、增强对象指派责任的灵活性
// 不能保证请求一定会被接受，调试不方便，可能有循环调用
type WordsFilter interface {
	Filter() bool
}

type WordsFilterChain struct {
	filters []WordsFilter
}

func (c *WordsFilterChain) AddFilter(wf WordsFilter) {
	c.filters = append(c.filters, wf)
}

func (c *WordsFilterChain) Check() bool {
	for _, wf := range c.filters {
		if !wf.Filter() {
			return false
		}
	}

	return true
}

type A struct {
}

func (a A) Filter() bool {
	fmt.Println("a is ok")
	return true
}

type B struct {
}

func (b B) Filter() bool {
	fmt.Println("b is failed")
	return false
}
