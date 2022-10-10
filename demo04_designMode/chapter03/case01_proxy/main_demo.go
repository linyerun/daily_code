// 代理模式：

package main

// Goods 货物
type Goods struct {
	Kind   string
	IsFake bool
}

// ---- 抽象层 ----

type Shopping interface {
	Buy(goods Goods)
}

// ---- 实现层 ----

type AmericanShopping struct {
}

func (t *AmericanShopping) Buy(goods Goods) {
	println("美国购物：[", "类型：", goods.Kind, "]")
}

type KoreanShopping struct {
}

func (t *KoreanShopping) Buy(goods Goods) {
	println("韩国购物：[", "类型：", goods.Kind, "]")
}

// ---- 代理层 ----

type proxy struct {
	shopping Shopping
}

func (t *proxy) ProxyShopping(goods Goods) {
	//1. 判断真伪
	if t.distinguish(goods) {
		//货物是假的
		return
	}
	//2. 进行购买
	t.shopping.Buy(goods)
	//3. 海关检验
	t.check(goods)
}

func (t *proxy) distinguish(goods Goods) bool {
	println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.IsFake {
		println("货物是假的")
	}
	return goods.IsFake
}

func (t *proxy) check(goods Goods) {
	println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func NewProxy(shopping Shopping) *proxy {
	return &proxy{shopping: shopping}
}

func main() {
	//韩国面膜
	KoreanProxy := NewProxy(new(KoreanShopping))
	KoreanProxy.ProxyShopping(Goods{Kind: "韩国面膜", IsFake: false})
	//代理美国
	AmericanProxy := NewProxy(new(AmericanShopping))
	AmericanProxy.ProxyShopping(Goods{Kind: "CET4", IsFake: true})
}
