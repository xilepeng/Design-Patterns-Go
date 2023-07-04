package main

import (
	"fmt"
)

type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品真伪
}

// 抽象层
type Shopping interface {
	Buy(goods *Goods) // 某任务
}

// 实现层
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国买了", goods.Kind)
}

// 美国购物
type AmericaShopping struct {
}

func (ks *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国买了", goods.Kind)
}

// 代理-海外代购
type OverSeasProxy struct {
	shopping Shopping // 代理某个主题，这里是抽象的数据类型
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	// 1. 辨别真伪
	if op.distinguish(goods) == true {
		// 2. 调用具体要被代理的购物方式的Buy()方法
		op.shopping.Buy(goods) // 多态
		// 3. 海关安检
		op.check(goods)
	}
}

// 辨别真伪的能力
func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, "不应该购买")
	}
	return goods.Fact
}

// 安检流程
func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查，成功带回祖国")
}

// 业务逻辑层

func main() {
	g1 := Goods{Kind: "韩国面膜", Fact: true}
	g2 := Goods{Kind: "CET4证书", Fact: false}

	/*
		var shopping Shopping
		shopping = new(KoreaShopping)

		// 1. 先辨别真伪
		if g1.Fact == true {
			// 2. 买
			shopping.Buy(g1)
			// 3. 带回祖国
			fmt.Println("带回祖国")
		}
	*/
	var kShopping Shopping
	kShopping = new(KoreaShopping)

	var proxy Shopping
	proxy = NewProxy(kShopping)
	proxy.Buy(&g1)

}
