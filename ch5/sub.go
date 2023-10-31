package ch5

import "fmt"

func Sub1() {
	mail := Mail{
		Type:        Normal,
		ToAddresses: []string{"recipient@example.com"},
		OwnAddress:  "sender@example.com",
		HasResMail:  false,
		HasFwMail:   false,
	}

	validators := []Validator{
		SharedValidator{},
		ResValidator{},
		FwValidator{},
	}

	result := Validate(mail, validators)
	fmt.Println(result)
}
