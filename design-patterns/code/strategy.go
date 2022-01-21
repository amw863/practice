package code

import "fmt"

/**
 * 策略模式：算法独立于业务
 * 完美支持开闭原则，高度灵活、提供了管理算法族的办法
 * 客户端必须知道所以策略
 */

type Strategy interface {
	Execute(p float32) float32
}

type NovelStrategy struct {
}

func (n NovelStrategy) Execute(p float32) float32 {
	return p * 0.1
}

type Book struct {
	price    float32
	disCount Strategy
}

func (b Book) Print() {
	fmt.Printf("原价  : %.2f \n折后价: %.2f\n", b.price, b.disCount.Execute(b.price))
}

func NewBook(price float32, s Strategy) Book {
	return Book{
		price: price, disCount: s,
	}
}
