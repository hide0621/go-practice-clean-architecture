package ch7

type SendMailUseCase interface {
	Execute(requestValue RequestValue) ResponseValue
}

type RequestValue struct {
	ToAddress string
	MainText  string
}

type ResponseValue struct {
	Result chan bool
}

type SendMailUseCaseImpl struct {
	MailGateway       MailGateway
	AccountRepository AccountRepository
}

func (useCase *SendMailUseCaseImpl) Execute(requestValue RequestValue) ResponseValue {
	ownAccount := useCase.AccountRepository.GetMainAccount()
	toAddress := requestValue.ToAddress
	result := useCase.MailGateway.SendMail(Mail{
		ownAddress: ownAccount.Address,
		toAddresses: &[]Address{{
			address: toAddress,
		}},
		mainText: requestValue.MainText,
	})
	return ResponseValue{Result: result}
}
