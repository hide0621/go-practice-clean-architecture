package ch7

type MailEditFragment struct {
}

type MailEditPresenter interface {
	LoadMail(mailID int)
}

type MailEditPresenterImpl struct {
	VM              *MailEditVM
	LoadMailUseCase LoadMailUseCase
}

func (presenter *MailEditPresenterImpl) LoadMail(mailID int) {
	mail := presenter.LoadMailUseCase.Execute(mailID)
	presenter.VM.SetMail(mail)
}

type MailEditVM struct {
}

func (vm *MailEditVM) SetMail(mail Mail) {

}

type MailEditModule struct {
}

func (module *MailEditModule) ProvideVM(fragment *MailEditFragment) *MailEditVM {

	return &MailEditVM{}
}
