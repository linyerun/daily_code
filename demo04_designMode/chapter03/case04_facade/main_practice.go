package main

import "fmt"

// TV 电视机
type TV struct{}

func (t *TV) On() {
	fmt.Println("打开 电视机")
}

func (t *TV) Off() {
	fmt.Println("关闭 电视机")
}

// VoiceBox 音箱
type VoiceBox struct{}

func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}

func (v *VoiceBox) Off() {
	fmt.Println("关闭 音箱")
}

// Light 灯光
type Light struct{}

func (l *Light) On() {
	fmt.Println("打开 灯光")
}

func (l *Light) Off() {
	fmt.Println("关闭 灯光")
}

// Xbox 游戏机
type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("打开 游戏机")
}

func (x *Xbox) Off() {
	fmt.Println("关闭 游戏机")
}

// MicroPhone 麦克风
type MicroPhone struct{}

func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}

// Projector 投影仪
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

func (p *Projector) Off() {
	fmt.Println("关闭 投影仪")
}

// homePlayerFacade 家庭影院(外观)
type homePlayerFacade struct {
	tv    *TV
	vb    *VoiceBox
	light *Light
	xbox  *Xbox
	mp    *MicroPhone
	pro   *Projector
}

func NewHomePlayerFacade(tv *TV, box *VoiceBox, light *Light, xbox *Xbox, phone *MicroPhone, projector *Projector) *homePlayerFacade {
	return &homePlayerFacade{tv: tv, vb: box, light: light, xbox: xbox, pro: projector, mp: phone}
}

// DoKTV KTV模式
func (hp *homePlayerFacade) DoKTV() {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.light.Off()
	hp.vb.On()
}

// DoGame 游戏模式
func (hp *homePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}

func main() {
	homePlayer := NewHomePlayerFacade(new(TV), new(VoiceBox), new(Light), new(Xbox), new(MicroPhone), new(Projector))

	homePlayer.DoKTV()

	fmt.Println("------------")

	homePlayer.DoGame()
}
