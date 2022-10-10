//装饰器模式：装饰模式比生成子类实现更为灵活

package main

// ---- 抽象层 ----

type Phone interface {
	Show()
}

// Decorator 由于golang语法局限性，下面这么写
type Decorator struct {
	phone Phone
}

func (t *Decorator) Show() {
	panic("你需要重写Decorator的Show方法")
}

// ---- 实现层 ----

type Huawei struct {
}

func (t *Huawei) Show() {
	println("秀出华为手机")
}

type Xiaomi struct {
}

func (t *Xiaomi) Show() {
	println("秀出你的小米手机")
}

type KeDecorator struct {
	Decorator
}

func (t *KeDecorator) Show() {
	t.phone.Show()
	println("给手机加个手机壳")
}

func NewKeDecorator(phone Phone) *KeDecorator {
	return &KeDecorator{Decorator: Decorator{phone: phone}}
}

type MoDecorator struct {
	Decorator
}

func (t *MoDecorator) Show() {
	t.phone.Show()
	println("给手机加个手机膜")
}

func NewMoDecorator(phone Phone) *MoDecorator {
	return &MoDecorator{Decorator: Decorator{phone: phone}}
}

// ---- 业务逻辑层 ----
func main() {
	huawei := new(Huawei)
	huawei.Show()

	xiaomi := new(Xiaomi)
	xiaomi.Show()

	NewKeDecorator(huawei).Show()
	NewMoDecorator(xiaomi).Show()

	NewMoDecorator(NewKeDecorator(xiaomi)).Show()
}
