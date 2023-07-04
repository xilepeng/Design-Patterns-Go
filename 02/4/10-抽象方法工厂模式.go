package main

import (
	"fmt"
)

// 抽象层

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层

// 中国产品簇

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("中国的苹果")
}

type ChinaBanana struct{}

func (ca *ChinaBanana) ShowBanana() {
	fmt.Println("中国的香蕉")
}

type ChinaPear struct{}

func (ca *ChinaPear) ShowPear() {
	fmt.Println("中国的梨")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

// 逻辑层
func main() {
	// 需求1. 需要中国的苹果，香蕉，梨

	// 1. 创建一个中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	// 2. 生成中国苹果
	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()

	// 3. 生成中国香蕉
	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()
	cBanana.ShowBanana()

	// 4. 生成中国梨
	var cPear AbstractPear
	cPear = cFac.CreatePear()
	cPear.ShowPear()

}
