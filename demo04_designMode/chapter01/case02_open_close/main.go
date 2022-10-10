//开闭原则: 类的改动是通过增加代码进行的，而不是修改源代码。

package main

import "fmt"

type AbstractBanker interface {
	Do()
}

// SaveBanker 存款的业务员
type SaveBanker struct{}

func (t *SaveBanker) Do() {
	fmt.Println("正在进行存款服务")
}

// TransferBanker 转账的业务员
type TransferBanker struct{}

func (t *TransferBanker) Do() {
	fmt.Println("正在进行转账服务")
}

// PaymentBanker 支付的业务员
type PaymentBanker struct{}

func (t *PaymentBanker) Do() {
	fmt.Println("正在进行支付服务")
}

// BankerDo 做一个小框架
func BankerDo(banker AbstractBanker) {
	banker.Do()
}

func main() {
	BankerDo(new(SaveBanker))
	BankerDo(new(TransferBanker))
	BankerDo(new(PaymentBanker))
}
