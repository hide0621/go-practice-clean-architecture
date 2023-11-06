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

	resProvider := ResProvider{value: true}
	soundProvider := SoundProvider{value: false}
	vibrationProvider := VibrationProvider{value: true}

	spm.AddBooleanProvider("res", resProvider)
	spm.AddBooleanProvider("sound", soundProvider)
	spm.AddBooleanProvider("vibration", vibrationProvider)

	fmt.Println("IsResPrefix: ", spm.sp["res"].IsEnabled())
	fmt.Println("IsEnableNotifySound:", spm.sp["sound"].IsEnabled())
	fmt.Println("IsEnableNotifyVib:", spm.sp["vibration"].IsEnabled())

}
