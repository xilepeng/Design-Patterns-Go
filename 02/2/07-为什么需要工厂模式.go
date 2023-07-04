package main

import "fmt"

// 抽象层
type Fruit interface {
	Show() // 接口的方法
}

// 实现层

type Apple struct {
	Fruit // 为了易于理解
}

func (apple *Apple) Show() {
	fmt.Printf("我是苹果\n")
}

type Pear struct {
	Fruit // 为了易于理解
}

func (pear *Pear) Show() {
	fmt.Printf("我是🍐\n")
}

type Banana struct {
	Fruit // 为了易于理解
}

func (banana *Banana) Show() {
	fmt.Printf("我是🍌\n")
}

// 工厂模块

type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		// apple 的初始化业务
		fruit = new(Apple) // 满足多态条件的赋值，父类指针指向子类对象
	} else if kind == "pear" {
		fruit = new(Pear)
	} else {
		fruit = new(Banana)
	}
	return fruit
}

// 业务逻辑层

func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()
}
