package main

import (
	"fmt"
	"sync"
)

// 懒汉式单例模式
type singelton struct{}

var instance *singelton

var once sync.Once

// 2.改进：原子操作
func GetInstance() *singelton {
	// 首次进入
	once.Do(func() {
		instance = new(singelton)
	})
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()
}
