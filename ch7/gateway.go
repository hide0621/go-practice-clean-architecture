package ch7

type MailGateway interface {
	SendMail(mail Mail) chan bool
	CheckNewMails() chan []Mail
}

type ApiClient struct {
}

type MailApiClient struct {
	ApiClient ApiClient
}

func (mac *MailApiClient) SendMail(mail Mail) chan bool {

	resultChan := make(chan bool)

	go func() {
		// 送信成功時
		// 送信失敗時はエラーを伴うように実装
		resultChan <- true
	}()

	return resultChan

}

func (mac *MailApiClient) CheckNewMails() chan []Mail {

	mailsChan := make(chan []Mail)

	go func() {
		// 取得成功時
		// 取得失敗時はエラーを伴うように実装
		newMails := []Mail{}
		mailsChan <- newMails
	}()

	return mailsChan
}
