//适配器模式：

package main

type V5 interface {
	UseV5()
}

// Phone 依赖V5进行充电
type Phone struct {
	v5 V5
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5: v5}
}

func (t *Phone) Charge() {
	println("手机正在充电.....")
	t.v5.UseV5()
}

// V220 存在一个220V的
type V220 struct {
}

func (t *V220) UseV220() {
	println("使用220v电压进行充电")
}

// Adapter 适配器
type Adapter struct {
	v220 *V220
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

func (t *Adapter) UseV5() {
	println("使用适配器进行充电")
	t.v220.UseV220()
}

func main() {
	NewPhone(NewAdapter(new(V220))).Charge()
}
