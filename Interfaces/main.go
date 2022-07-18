package main

import "fmt"

type Debugger interface {
	Debug() string
}

type CalculatorIface interface{
	Calc() int
}

type Calculator struct {
	A int
	B int
}

func (c Calculator) Calc() int {
	return c.A + c.B
}

//https://cheatography.com/fenistil/cheat-sheets/go-fmt-formattings/
func (c Calculator) Debug() string {
	return fmt.Sprintf("A = %v, B = %v", c.A, c.B)
}



type Calculator2 struct {
	A int
	B int
}

func (c Calculator2) Calc() int {
	return c.A + c.B
}

//https://cheatography.com/fenistil/cheat-sheets/go-fmt-formattings/
func (c Calculator2) Debug() string {
	return fmt.Sprintf("A = %v, B = %v", c.A, c.B)
}

func Log(i Debugger) {
	fmt.Println("Log:", i.Debug())
}

func main() {
	calc := Calculator{
		A: 1,
		B: 2,


	calc2 := Calculator2{
		A: 1,
		B: 2,
	}
	fmt.Println(calc.Calc())
	fmt.Println(calc.Debug())
	Log(calc)

}
