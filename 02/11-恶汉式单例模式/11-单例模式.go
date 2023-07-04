package main

import "fmt"

// 恶汉式单例模式
type singelton struct{}

var instance *singelton = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()
}
