package main

import (
	"fmt"
	"go-practice-clean-architecture/ch4"
	"go-practice-clean-architecture/ch5"
)

func main() {

	buttonInfo1 := ch4.BtnInfo{X: ch4.ButtonState{IsPress: true}, Y: ch4.ButtonState{IsPress: true}}
	ch4.ProcessA(buttonInfo1)

	buttonInfo2 := ch4.BtnInfo{X: ch4.ButtonState{IsPress: false}, Y: ch4.ButtonState{IsPress: false}}
	ch4.ProcessB(buttonInfo2)

	mail := ch5.Mail{
		Type:        ch5.Normal,
		ToAddresses: []string{"recipient@example.com"},
		OwnAddress:  "sender@example.com",
		HasResMail:  false,
		HasFwMail:   false,
	}
	result := ch5.Validate(mail)
	fmt.Println(result)

}
