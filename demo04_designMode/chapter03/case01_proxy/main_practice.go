// 潘金莲、西门庆、王婆 => 这个例子纯属搞笑，对刘丹冰的例子进行了修改(感觉这修改不好)

package main

// ---- 抽象层 ----

type BeautyWoman interface {
	MakeEyesWithMan()
	HappyWithMan()
}

type Man struct {
	Name string
}

// ---- 实现层 ----

type PanJinLian struct {
	man *Man
}

func (t *PanJinLian) NewPanJinLian(man *Man) *PanJinLian {
	return &PanJinLian{man: man}
}

func (t *PanJinLian) MakeEyesWithMan() {
	println("潘金莲对", t.man.Name, "抛媚眼")
}

func (t *PanJinLian) HappyWithMan() {
	println("潘金莲和", t.man.Name, "浪漫约会中......")
}

// ---- 代理类 ----

type WangPo struct {
	woman BeautyWoman
}

func NewWangPo(woman BeautyWoman) *WangPo {
	return &WangPo{woman: woman}
}

func (t *WangPo) MakeEyesWithMan() {
	t.woman.MakeEyesWithMan()
}

func (t *WangPo) HappyWithMan() {
	t.woman.HappyWithMan()
}

// 业务逻辑流
func main() {
	//西门庆
	xiMenQing := &Man{Name: "西门庆"}
	//西门庆找王婆约潘金莲
	wangPo := NewWangPo(&PanJinLian{xiMenQing})
	wangPo.MakeEyesWithMan()
	wangPo.HappyWithMan()
}
