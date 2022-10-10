//抽象工厂模式：它适用于产品等级趋于稳定，需要增加产品族；它对产品族是符合开闭原则的，但是对于产品等级就不符合了
//什么是产品等级、产品族：看下面例子，获取刘丹冰大佬的笔记

package main

// ---- 抽象层 ----

type AbstractApple interface {
	ShowAbstractApple()
}
type AbstractBanana interface {
	ShowAbstractBanana()
}
type AbstractPear interface {
	ShowAbstractPear()
}
type AbstractFactory interface {
	CreateAbstractApple() AbstractApple
	CreateAbstractBanana() AbstractBanana
	CreateAbstractPear() AbstractPear
}

// ---- 实现层 ----

//中国族的水果

type ChineseApple struct {
}

func (t *ChineseApple) ShowAbstractApple() {
	println("中国苹果")
}

type ChineseBanana struct {
}

func (t *ChineseBanana) ShowAbstractBanana() {
	println("中国香蕉")
}

type ChinesePear struct {
}

func (t *ChinesePear) ShowAbstractPear() {
	println("中国鸭梨")
}

type ChineseAbstractFactory struct {
}

func (t ChineseAbstractFactory) CreateAbstractApple() AbstractApple {
	return new(ChineseApple)
}
func (t ChineseAbstractFactory) CreateAbstractBanana() AbstractBanana {
	return new(ChineseBanana)
}
func (t ChineseAbstractFactory) CreateAbstractPear() AbstractPear {
	return new(ChinesePear)
}

//美国族的水果

type AmericanApple struct {
}

func (t *AmericanApple) ShowAbstractApple() {
	println("美国苹果")
}

type AmericanBanana struct {
}

func (t *AmericanBanana) ShowAbstractBanana() {
	println("美国香蕉")
}

type AmericanPear struct {
}

func (t *AmericanPear) ShowAbstractPear() {
	println("美国鸭梨")
}

type AmericanAbstractFactory struct {
}

func (t *AmericanAbstractFactory) CreateAbstractApple() AbstractApple {
	return new(AmericanApple)
}
func (t *AmericanAbstractFactory) CreateAbstractBanana() AbstractBanana {
	return new(AmericanBanana)
}
func (t *AmericanAbstractFactory) CreateAbstractPear() AbstractPear {
	return new(AmericanPear)
}

func main() {
	// 中国的
	ca := new(ChineseAbstractFactory)
	ca.CreateAbstractApple().ShowAbstractApple()
	ca.CreateAbstractBanana().ShowAbstractBanana()
	ca.CreateAbstractPear().ShowAbstractPear()
	// 美国的
	aa := new(AmericanAbstractFactory)
	aa.CreateAbstractApple().ShowAbstractApple()
	aa.CreateAbstractBanana().ShowAbstractBanana()
	aa.CreateAbstractPear().ShowAbstractPear()
}
