package main

import "fmt"

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Printf("Benz is running ... \n\n")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Printf("BMW is running ... \n\n")
}

type ZhangSan struct{}

func (z *ZhangSan) Drive(car Car) {
	fmt.Println("张三 Drive car ...")
	car.Run()
}

type LiSi struct{}

func (z *LiSi) Drive(car Car) {
	fmt.Println("李四 Drive car ...")
	car.Run()
}

// 业务逻辑层

func main() {

	// 张三开奔驰
	var benz Car
	benz = new(Benz)
	var zhangSan Driver
	zhangSan = new(ZhangSan)
	zhangSan.Drive(benz)

	// 李四开宝马
	var bmw Car
	bmw = new(BMW)
	var lisi Driver
	lisi = new(LiSi)
	lisi.Drive(bmw)

	// 王五开奔驰

}
