package ch7

type MailRepository interface {
	LoadMail(mailID int) Mail
	SaveMail(ownAddress string, mailID int, toAddress string, mainText string)
}

type RoomDb struct {
}

type RoomManager struct {
	DB RoomDb
}

func (rm *RoomManager) LoadMail(mailID int) Mail {
	return Mail{}
}

func (rm *RoomManager) SaveMail(ownAddress *Address, mailID int, toAddress *[]Address, mainText string) {

}
