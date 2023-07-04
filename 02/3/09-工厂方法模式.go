package main

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

// 水果类（抽象的接口）

// 工厂类（抽象的接口）

// 基础模块层

// 实现层

type Apple struct {
	Fruit // 为了易于理解
}

func (apple *Apple) Show() {
	fmt.Printf("我是🍎\n")
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

// 基础的工厂模块
type AbstractFactory interface {
	CreateFruit() Fruit // 生成水果类，抽象的生产器方法
}

// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的苹果
	fruit = new(Apple)
	return fruit
}

// 具体的梨工厂

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的梨
	fruit = new(Pear)
	return fruit
}

// 具体的香蕉工厂

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	// 生产一个具体的香蕉
	fruit = new(Banana)
	return fruit
}

func main() {
	// 1. 需要一个具体的苹果对象
	// 0. 先要一个具体的苹果工厂
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	// 2. 生成一个具体的水果
	var apple Fruit
	apple = appleFactory.CreateFruit()
	apple.Show() // 多态

	// 2. 需要一个具体的香蕉对象
	var bananaFactory AbstractFactory
	bananaFactory = new(BananaFactory)
	var banana Fruit
	banana = bananaFactory.CreateFruit()
	banana.Show()

	// 3. 需要一个具体的梨对象
	pearFactory := new(PearFactory)
	pear := pearFactory.CreateFruit()
	pear.Show()
}
