package ch7

import "fmt"

func Sub1() {

	ownAddr, err := NewAddress("example@example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	toAddrs := &[]Address{
		{
			address: "recipient1@example.com",
		},
		{
			address: "recipient2@example.com",
		},
	}

	mail := Mail{
		id:          1,
		ownAddress:  ownAddr,
		toAddresses: toAddrs,
		mainText:    "Hello, this is the main text of the mail.",
	}

	fmt.Println(mail)
}
