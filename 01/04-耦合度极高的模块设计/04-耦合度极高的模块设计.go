package main

import "fmt"

// 奔驰汽车
type Benz struct{}

// 奔驰汽车的方法
func (b *Benz) Run() {
	fmt.Printf("Benz is running ...\n\n")
}

// 宝马汽车
type BMW struct{}

// 宝马汽车的方法
func (b *BMW) Run() {
	fmt.Printf("BMW is running ...\n\n")
}

// 司机张三
type ZhangSan struct {
}

func (z *ZhangSan) DriveBenz(benz *Benz) {
	fmt.Println("张三开奔驰...")
	benz.Run()
}

// 司机李四
type Lisi struct{}

func (l *Lisi) DriveBMw(bmw *BMW) {
	fmt.Println("李四开宝马...")
	bmw.Run()
}
func main() {
	benz, zhangSan := &Benz{}, &ZhangSan{}
	zhangSan.DriveBenz(benz)

	bmw, lisi := &BMW{}, &Lisi{}
	lisi.DriveBMw(bmw)
}
