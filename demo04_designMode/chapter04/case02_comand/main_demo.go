// 命令模式：

package main

type Doctor struct {
}

func (d *Doctor) TreatEyes() {
	println("医生治疗眼睛")
}

func (d *Doctor) TreatNose() {
	println("医生治疗鼻子")
}

type Command interface {
	Treat()
}

//治疗眼睛的指令
type cmdTreatEyes struct {
	d *Doctor
}

func NewCmdTreatEyes(doctor *Doctor) *cmdTreatEyes {
	return &cmdTreatEyes{d: doctor}
}

func (t *cmdTreatEyes) Treat() {
	t.d.TreatEyes()
}

//治疗鼻子的指令
type cmdTreatNose struct {
	d *Doctor
}

func NewCmdTreatNose(doctor *Doctor) *cmdTreatEyes {
	return &cmdTreatEyes{d: doctor}
}

func (t *cmdTreatNose) Treat() {
	t.d.TreatNose()
}

// Nurse 护士：负责收集病人要指令的指令，然后发送给医生
type Nurse struct {
	cmdList []Command
}

func (t *Nurse) AddCmds(cmds ...Command) {
	t.cmdList = append(t.cmdList, cmds...)
}

func (t *Nurse) Notify() {
	for _, cmd := range t.cmdList {
		cmd.Treat()
	}
}

func main() {
	cmdTreatNose := NewCmdTreatNose(new(Doctor))
	cmdTreatEyes := NewCmdTreatEyes(new(Doctor))

	nurse := new(Nurse)
	nurse.AddCmds(cmdTreatNose, cmdTreatEyes)

	nurse.Notify()
}
