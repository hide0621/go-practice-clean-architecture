package ch7

import "fmt"

func Sub1() {

	ownAddr := NewAddress("example@example.com")

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

func Sub2() {

	requestValue := RequestValue{
		ToAddress: "recipient1@example.com",
		MainText:  "Hello, this is the main text of the mail.",
	}

	repository := AccountRepositoryImpl{}

	gateway := MailGatewayImpl{}

	useCase := SendMailUseCaseImpl{
		MailGateway:       &gateway,
		AccountRepository: &repository,
	}

	useCase.Execute(requestValue)

}
