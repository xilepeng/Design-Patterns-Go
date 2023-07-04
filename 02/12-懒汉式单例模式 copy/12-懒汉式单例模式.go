package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 懒汉式单例模式
type singelton struct{}

var instance *singelton

// 定义一个互斥锁
var lock sync.Mutex

// 标记
var initialized uint32

// 1.加锁太慢
// func GetInstance() *singelton {
// 	// 为了线程安全，增加互斥锁
// 	lock.Lock()
// 	defer lock.Unlock()

// 	// 只有首次GetInstance() 方法被调用，才会生成这个到单例的对象
// 	if instance == nil {
// 		instance = new(singelton)
// 		return instance
// 	}
// 	return instance
// }

// 2.改进：原子操作
func GetInstance() *singelton {
	// 首次进入
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	// 如果没有，加锁申请
	lock.Lock()
	defer lock.Unlock()

	// 只有首次GetInstance() 方法被调用，才会生成这个到单例的对象
	if instance == nil {
		instance = new(singelton)
		// 设置标记为1
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()
}
