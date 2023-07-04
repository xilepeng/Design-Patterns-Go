package main

import "fmt"

type Clothes struct{}

// 工作穿衣服
func (c *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}

// 逛街的装扮
func (c *Clothes) OnShop() {
	fmt.Println("逛街的装扮")
}
func main() {
	c := Clothes{}
	// 工作的业务
	fmt.Println("在工作...")
	c.OnWork()
	fmt.Println("")
	// 逛街的业务
	fmt.Println("在逛街...")
	c.OnShop()
}
