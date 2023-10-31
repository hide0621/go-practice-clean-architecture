package ch5

type MailType int

const (
	Normal MailType = iota
	Res
	Fw
)

type Validator interface {
	Validate(mail Mail) bool
}

type Mail struct {
	Type        MailType
	ToAddresses []string
	OwnAddress  string
	HasResMail  bool
	HasFwMail   bool
}

const Empty = ""

type SharedValidator struct{}

func (sv SharedValidator) Validate(mail Mail) bool {
	return len(mail.ToAddresses) > 0 && mail.OwnAddress != Empty
}

type ResValidator struct{}

func (rv ResValidator) Validate(mail Mail) bool {
	if mail.Type != Res {
		return true
	}
	return mail.HasResMail
}

type FwValidator struct{}

func (fv FwValidator) Validate(mail Mail) bool {
	if mail.Type != Fw {
		return true
	}
	return mail.HasFwMail
}

func Validate(mail Mail, validators []Validator) bool {
	result := true
	for _, validator := range validators {
		result = result && validator.Validate(mail)
	}
	return result
}
