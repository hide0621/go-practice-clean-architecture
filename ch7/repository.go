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

type AccountRepository interface {
	GetMainAccount() *Account
}

type Account struct {
	Address *Address
}

type AccountRepositoryImpl struct {
	// アカウントリポジトリの具体的な実装
}

func (repository *AccountRepositoryImpl) GetMainAccount() *Account {
	// メインアカウントの取得ロジックの実装
	return &Account{
		Address: NewAddress("main@example.com"),
		// 他のアカウント関連のプロパティ
	}
}
