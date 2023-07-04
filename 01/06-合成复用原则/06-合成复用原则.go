package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Printf("小猫吃饭...\n")
}

// 1.使用继承添加睡觉方法
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Printf("（继承）小猫睡觉...\n")
}

// 2.使用组合添加睡觉方法
type CatC struct {
	C *Cat // 组合进来一个 Cat 类
}

func (cc *CatC) Eat(c *Cat) {
	// cc.C.Eat() // 调用组合进来的原有类的功能
	c.Eat() // 调用组合进来的原有类的功能
}

func (cc *CatC) Sleep2() {
	fmt.Printf("(组合）小猫睡觉...")
}
func main() {
	var cat Cat
	cat.Eat()

	fmt.Printf("1.使用继承添加睡觉方法\n")
	var catb CatB
	catb.Eat()
	catb.Sleep()

	fmt.Printf("2.使用组合添加睡觉方法\n")
	var catc CatC
	catc.Eat(&cat)
	catc.Sleep2()
}
