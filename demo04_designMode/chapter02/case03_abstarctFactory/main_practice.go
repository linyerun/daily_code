package main

// ---- 抽象层 ----

type GPU interface {
	Display()
}
type MemoryBank interface {
	Storage()
}
type CPU interface {
	Calculate()
}

type AbstractComputerFactory interface {
	CreateGPU() GPU
	CreateMemoryBank() MemoryBank
	CreateCPU() CPU
}

// ---- 实现层 ----

// Inter

type InterGPU struct {
}

func (t *InterGPU) Display() {
	println("Inter生产的显卡：开始display")
}

type InterMemoryBank struct {
}

func (t *InterMemoryBank) Storage() {
	println("Inter生产的内存：开始storage")
}

type InterCPU struct {
}

func (t *InterCPU) Calculate() {
	println("Inter生产的CPU：开始calculate")
}

type InterFactory struct {
}

func (t *InterFactory) CreateGPU() GPU {
	return new(InterGPU)
}
func (t *InterFactory) CreateMemoryBank() MemoryBank {
	return new(InterMemoryBank)
}
func (t *InterFactory) CreateCPU() CPU {
	return new(InterCPU)
}

// Nvidia

type NvidiaGPU struct {
}

func (t *NvidiaGPU) Display() {
	println("Nvidia生产的显卡：开始display")
}

type NvidiaMemoryBank struct {
}

func (t *NvidiaMemoryBank) Storage() {
	println("Nvidia生产的内存：开始storage")
}

type NvidiaCPU struct {
}

func (t *NvidiaCPU) Calculate() {
	println("Nvidia生产的CPU：开始calculate")
}

type NvidiaFactory struct {
}

func (t *NvidiaFactory) CreateGPU() GPU {
	return new(NvidiaGPU)
}
func (t *NvidiaFactory) CreateMemoryBank() MemoryBank {
	return new(NvidiaMemoryBank)
}
func (t *NvidiaFactory) CreateCPU() CPU {
	return new(NvidiaCPU)
}

// Kingston

type KingstonGPU struct {
}

func (t *KingstonGPU) Display() {
	println("Kingston生产的显卡：开始display")
}

type KingstonMemoryBank struct {
}

func (t *KingstonMemoryBank) Storage() {
	println("Kingston生产的内存：开始storage")
}

type KingstonCPU struct {
}

func (t *KingstonCPU) Calculate() {
	println("Kingston生产的CPU：开始calculate")
}

type KingstonFactory struct {
}

func (t *KingstonFactory) CreateGPU() GPU {
	return new(KingstonGPU)
}
func (t *KingstonFactory) CreateMemoryBank() MemoryBank {
	return new(KingstonMemoryBank)
}
func (t *KingstonFactory) CreateCPU() CPU {
	return new(KingstonCPU)
}

//组装电脑

type Computer struct {
	gpu        GPU
	memoryBank MemoryBank
	cpu        CPU
}

func (c *Computer) ShowComputer() {
	println("----", "电脑配置", "----")
	c.gpu.Display()
	c.memoryBank.Storage()
	c.cpu.Calculate()
	println()
}

func NewComputer(gpu GPU, memoryBank MemoryBank, cpu CPU) *Computer {
	return &Computer{gpu: gpu, memoryBank: memoryBank, cpu: cpu}
}

//业务逻辑层
func main() {
	interFac := new(InterFactory)
	nvidiaFac := new(NvidiaFactory)
	kingstonFac := new(KingstonFactory)
	//第一台电脑
	NewComputer(interFac.CreateGPU(), interFac.CreateMemoryBank(), interFac.CreateCPU()).ShowComputer()
	//第二台电脑
	NewComputer(nvidiaFac.CreateGPU(), kingstonFac.CreateMemoryBank(), interFac.CreateCPU()).ShowComputer()
}
