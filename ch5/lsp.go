package ch5

import "fmt"

type MailType int

const (
	Normal MailType = iota
	Res
	Fw
)

type Mail struct {
	Type        MailType
	ToAddresses []string
	OwnAddress  string
	HasResMail  bool
	HasFwMail   bool
}

const Empty = ""

func Validate(mail Mail) bool {
	switch mail.Type {
	case Normal:
		return len(mail.ToAddresses) > 0 && mail.OwnAddress != Empty
	case Res:
		return len(mail.ToAddresses) > 0 && mail.OwnAddress != Empty && mail.HasResMail
	case Fw:
		return len(mail.ToAddresses) > 0 && mail.OwnAddress != Empty && mail.HasFwMail
	default:
		fmt.Println("Unsupported MailType")
		return false
	}
}
