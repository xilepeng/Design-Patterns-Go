package main

import "fmt"

// 抽象的银行业务员
type AbstractBanker interface {
	Do() // 抽象的处理业务接口
}

// 存款的业务员
type SaveBanker struct {
}

func (sb *SaveBanker) Do() {
	fmt.Println("进行了存款\n")
}

// 转账的业务员
type TransferBanker struct{}

func (tb *TransferBanker) Do() {
	fmt.Println("进行了转账")
}

// 实现一个架构层（基于抽象层进行业务封装，针对interface接口进行封装（
func BankBusiness(bank AbstractBanker) {
	// 通过接口向下调用（多态现象）
	bank.Do()
}
func main() {

	/*
		// 存款的业务
		sb := SaveBanker{}
		sb.Do()

		// 转账的业务
		tb := TransferBanker{}
		tb.Do()
	*/

	BankBusiness(&SaveBanker{})
	BankBusiness(&TransferBanker{})

}
