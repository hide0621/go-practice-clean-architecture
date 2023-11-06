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

func Sub2() {

	spm := NewSharedPreferenceManager()

	spm.SetBoolean("res", true)
	spm.SetBoolean("sound", false)
	spm.SetBoolean("vibration", true)

	fmt.Println("IsResPrefix: ", spm.IsResPrefix())
	fmt.Println("IsEnableNotifySound:", spm.IsEnableNotifySound())
	fmt.Println("IsEnableNotifyVib:", spm.IsEnableNotifyVib())

}
