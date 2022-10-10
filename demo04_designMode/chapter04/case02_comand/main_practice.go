package main

type chef struct {
	Name string
}

func NewChef(Name string) *chef {
	return &chef{Name: Name}
}

func (t *chef) ToastMutton() {
	println(t.Name, "开始烤羊肉串")
}

func (t *chef) ToastChickenWing() {
	println(t.Name, "开始烤鸡翅")
}

// Cmd 抽象指令
type Cmd interface {
	Toast()
}

//烤羊肉命令
type cmdToastMutton struct {
	c *chef
}

func NewCmdToastMutton(c *chef) *cmdToastMutton {
	return &cmdToastMutton{c: c}
}

func (t *cmdToastMutton) Toast() {
	t.c.ToastMutton()
}

//烤鸡翅命令
type cmdToastChickenWing struct {
	c *chef
}

func NewCmdToastChickenWing(c *chef) *cmdToastChickenWing {
	return &cmdToastChickenWing{c: c}
}

func (t *cmdToastChickenWing) Toast() {
	t.c.ToastChickenWing()
}

// Waiter 服务员
type Waiter struct {
	cmds []Cmd
}

func (t *Waiter) AddCmds(cmds ...Cmd) {
	t.cmds = append(t.cmds, cmds...)
}

func (t *Waiter) Notify() {
	for _, cmd := range t.cmds {
		cmd.Toast()
	}
}

func main() {
	cmd01 := NewCmdToastChickenWing(NewChef("王师傅"))
	cmd02 := NewCmdToastMutton(NewChef("林师傅"))
	cmd03 := NewCmdToastChickenWing(NewChef("张师傅"))

	waiter := new(Waiter)
	waiter.AddCmds(cmd01, cmd02, cmd03)

	waiter.Notify()
}
