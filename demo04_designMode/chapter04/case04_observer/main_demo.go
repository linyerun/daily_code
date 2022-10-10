// 观察者模式：

package main

// ---- 抽象层 ----

type Observer interface {
	OnTeacherComing()
}

type Notifier interface {
	AddObservers(observers ...Observer)
	RemoveObserver(observer Observer)
	Notify()
}

// ---- 实现层 ----

type ZhangSan struct {
}

func (t *ZhangSan) OnTeacherComing() {
	println("老师来了，张三停止打王者农药")
}

func (t *ZhangSan) DoBadThing() {
	println("张三正在打王者农药")
}

type LiSi struct {
}

func (t *LiSi) OnTeacherComing() {
	println("老师来了，李四停止抄作业")
}

func (t *LiSi) DoBadThing() {
	println("李四正在炒别人的作业")
}

type WangWu struct {
}

func (t *WangWu) OnTeacherComing() {
	println("老师来了，王五停止看别人打王者农药")
}

func (t *WangWu) DoBadThing() {
	println("王五正在看别人打王者农药")
}

type Monitor struct {
	observers []Observer
}

func (t *Monitor) AddObservers(observers ...Observer) {
	t.observers = append(t.observers, observers...)
}

func (t *Monitor) RemoveObserver(observer Observer) {
	for i := range t.observers {
		// 直接使用==，我们知道这实际在比较对象地址
		if t.observers[i] == observer {
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			break
		}
	}
}

func (t *Monitor) Notify() {
	for _, observer := range t.observers {
		observer.OnTeacherComing()
	}
}

// ---- 业务逻辑层 ----
func main() {
	zs := new(ZhangSan)
	ls := new(LiSi)
	ww := new(WangWu)
	monitor := new(Monitor)

	monitor.AddObservers(zs, ls, ww)

	zs.DoBadThing()
	ls.DoBadThing()
	ww.DoBadThing()

	monitor.Notify()

	monitor.RemoveObserver(zs)

	monitor.Notify()
}
