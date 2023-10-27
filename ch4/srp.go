package ch4

import "fmt"

type BtnInfo struct {
	X, Y ButtonState
}

type ButtonState struct {
	IsPress bool
}

func funcA1() {
	fmt.Println("Function A1 is called.")
}

func funcA2() {
	fmt.Println("Function A2 is called.")
}

func funcB1() {
	fmt.Println("Function B1 is called.")
}

func funcB2() {
	fmt.Println("Function B2 is called.")
}

func ProcessA(btn BtnInfo) {
	if btn.X.IsPress {
		if btn.Y.IsPress {
			funcA1()
		} else {
			funcA2()
		}
	} else {
		fmt.Println("X button is not pressed.")
	}
}

func ProcessB(btn BtnInfo) {
	if btn.X.IsPress {
		if btn.Y.IsPress {
			funcB1()
		} else {
			funcB2()
		}
	} else {
		fmt.Println("X button is not pressed.")
	}
}
