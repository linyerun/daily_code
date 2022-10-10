// 江湖混杂版观察者模式例子
// 丐帮：黄蓉、洪七公、乔峰
// 明教：张无忌、灭绝师太、金毛狮王
// 江湖百晓生：负责广播消息

package main

import "fmt"

const (
	GaiBang  string = "丐帮"
	MingJiao string = "明教"
)

// ---- 抽象层 ----

type Observer interface {
	OnFriendBeFight(event *Event)
	Title() string
	GetName() string
	GetParty() string
}

type Notifier interface {
	AddObservers(observers ...Observer)
	RemoveObserver(observer Observer)
	Notify(event *Event)
}

type Event struct {
	One     Observer
	Another Observer
	Msg     string
}

// ---- 实现层 ----

type Hero struct {
	Name  string
	Party string
}

func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]:%s", hero.Party, hero.Name)
}

func (hero *Hero) Fight(another Observer, baiXiaoSheng Notifier) {
	//构造事件
	e := new(Event)
	e.One = hero
	e.Another = another
	e.Msg = fmt.Sprintf("%s 将 %s 揍了...", hero.Title(), another.Title())

	//百晓生通知
	baiXiaoSheng.Notify(e)
}

func (hero *Hero) OnFriendBeFight(event *Event) {
	//判断是否为当事人
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}
	//本帮派同伴将其他门派揍了，要拍手叫好!
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Title(), "得知消息，拍手叫好！！！")
		return
	}

	//本帮派同伴被其他门派揍了，要主动报仇反击!
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Title(), "得知消息，发起报仇反击！！！")
		return
	}
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetParty() string {
	return hero.Party
}

type BaiXiaoSheng struct {
	heroList []Observer
}

func (b *BaiXiaoSheng) AddObservers(listeners ...Observer) {
	b.heroList = append(b.heroList, listeners...)
}

func (b *BaiXiaoSheng) RemoveObserver(listener Observer) {
	for index, l := range b.heroList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			b.heroList = append(b.heroList[:index], b.heroList[index+1:]...)
			break
		}
	}
}

func (b *BaiXiaoSheng) Notify(event *Event) {
	fmt.Println("【世界消息】 百晓生广播消息: ", event.Msg)
	for _, listener := range b.heroList {
		//依次调用全部观察的具体动作
		listener.OnFriendBeFight(event)
	}
}

// ---- 业务逻辑层 ----
func main() {
	hero1 := &Hero{
		"黄蓉",
		GaiBang,
	}

	hero2 := &Hero{
		"洪七公",
		GaiBang,
	}

	hero3 := &Hero{
		"乔峰",
		GaiBang,
	}

	hero4 := &Hero{
		"张无忌",
		MingJiao,
	}

	hero5 := &Hero{
		"韦一笑",
		MingJiao,
	}

	hero6 := &Hero{
		"金毛狮王",
		MingJiao,
	}

	baiXiaoSheng := BaiXiaoSheng{}
	baiXiaoSheng.AddObservers(hero1, hero2, hero3, hero4, hero5, hero6)
	fmt.Println("武林一片平静.....")
	hero1.Fight(hero5, &baiXiaoSheng)
}
