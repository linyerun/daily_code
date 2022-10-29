package main

type TestCombination struct {
}

func (t TestCombination) Func01() {

}

func (t *TestCombination) Func02() {

}

func main() {
	TestCombination.Func02(TestCombination{})
	(*TestCombination).Func01(new(TestCombination))
	(*TestCombination).Func02(new(TestCombination))
}
