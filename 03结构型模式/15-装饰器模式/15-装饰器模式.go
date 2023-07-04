package main

import "fmt"

// 抽象层
type Phone interface {
	Show() //构件的功能
}

// 抽象的装饰器，装饰器的基础类，该类本应该是 interface
// 但 Golang 的 interface 语法不可以有成员属性
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// 实现层

// 具体的构建
type Apple struct{}

func (hw *Apple) Show() {
	type Apple struct{}
	fmt.Println("秀出了 Apple 手机")
}

type Huawei struct{}

func (hw *Huawei) Show() {
	fmt.Println("秀出了HUAWEI手机")
}

// 具体的装饰器
type MoDecorator struct {
	Decorator // 继承基础的装饰器类（主要继承phone的成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show()      // 调用被装饰的构件的原方法
	fmt.Println("贴膜的手机") // 装饰器额外装饰的功能
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

// 业务逻辑层

func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show()
}
