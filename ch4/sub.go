package ch4

func Sub() {
	buttonInfo1 := BtnInfo{X: ButtonState{IsPress: true}, Y: ButtonState{IsPress: true}}
	ProcessA(buttonInfo1)

	buttonInfo2 := BtnInfo{X: ButtonState{IsPress: false}, Y: ButtonState{IsPress: false}}
	ProcessB(buttonInfo2)
}
