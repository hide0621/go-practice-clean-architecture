package main

import (
	"go-practice-clean-architecture/ch4"
)

func main() {

	buttonInfo1 := ch4.BtnInfo{X: ch4.ButtonState{IsPress: true}, Y: ch4.ButtonState{IsPress: true}}
	ch4.ProcessA(buttonInfo1)

	buttonInfo2 := ch4.BtnInfo{X: ch4.ButtonState{IsPress: false}, Y: ch4.ButtonState{IsPress: false}}
	ch4.ProcessB(buttonInfo2)

}