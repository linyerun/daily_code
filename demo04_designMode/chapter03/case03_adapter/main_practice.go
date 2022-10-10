// 适配器例子

package main

type Attack interface {
	Fight()
}

//技能

type DaBaoJian struct {
}

func (t *DaBaoJian) Fight() {
	println("使用大保健进行攻击")
}

type hero struct {
	name   string
	attack Attack
}

func NewHero(Name string, attack Attack) *hero {
	return &hero{name: Name, attack: attack}
}

func (t *hero) Kill() {
	println(t.name, "开始发动攻击")
	t.attack.Fight()
}

type Shutdown struct {
}

func (t *Shutdown) Shutdown() {
	println("关闭电脑")
}

//要求：盖伦的技能是Shutdown

type shutdownAdapter struct {
	shutdown *Shutdown
}

func NewShutdownAdapter(shutdown *Shutdown) *shutdownAdapter {
	return &shutdownAdapter{shutdown: shutdown}
}

func (t *shutdownAdapter) Fight() {
	t.shutdown.Shutdown()
}

func main() {
	NewHero("盖伦", new(DaBaoJian)).Kill()

	NewHero("盖伦", NewShutdownAdapter(new(Shutdown))).Kill()
}
