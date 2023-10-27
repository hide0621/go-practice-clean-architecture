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

func Process(btn BtnInfo, productType string) {
	if btn.X.IsPress {
		if btn.Y.IsPress {
			if productType == "A" {
				funcA1()
			} else if productType == "B" {
				funcB1()
			}
		} else {
			if productType == "A" {
				funcA2()
			} else if productType == "B" {
				funcB2()
			}
		}
	} else {
		fmt.Println("X button is not pressed.")
	}
}
