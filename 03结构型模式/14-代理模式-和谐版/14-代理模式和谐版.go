package main

import "fmt"

// 1. 抽象主题
type BeautyMoman interface {
	// 对男人抛媚眼
	MakeEyesWithMan()
	// 和男人共度美好时光
	HappyWithMan()
}

// 2. 具体主题
type PanJinlian struct {
}

// 对男人抛媚眼
func (p *PanJinlian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛了媚眼")
}

// 和男人共度美好时光
func (p *PanJinlian) HappyWithMan() {
	fmt.Println("潘金莲和本官共度了约会")
}

// 中间代理人，王婆
type WangPo struct {
	woman BeautyMoman
}

func NewProxy(woman BeautyMoman) BeautyMoman {
	return &WangPo{woman: woman}
}

// 王婆对男人抛媚眼
func (w *WangPo) MakeEyesWithMan() {
	fmt.Println("王婆对本官抛了媚眼")
}

// 王婆和男人共度美好时光
func (w *WangPo) HappyWithMan() {
	fmt.Println("潘金莲和本官共度了约会")
}

// 3. main 业务逻辑 - 西门大官人

func main() {

	// 大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinlian))
	// 王婆命令金莲抛媚眼
	wangpo.MakeEyesWithMan()
	// 王婆命令金莲和西门大官人约会
	wangpo.HappyWithMan()
}
