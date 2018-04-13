package validator

import "demo/modules/common/base"

type MobileValidator struct {
	base.Validator
}

func NewMobileValidator() *base.Validator {
	MobileValidator := &MobileValidator{}
	MobileValidator.Sence = map[string][]string{
		"mobileCode": []string{"phone", "smsCode"},
	}

	MobileValidator.Messages = map[string][]string{
		"phone":  []string{"required:rMobile", "digits:rDigits"},
		"smsCode": []string{"required:rSmsCode"},
	}

	MobileValidator.Rules = map[string][]string{
		"phone":  []string{"required", "digits:11"},
		"smsCode": []string{"required"},
	}
	return &MobileValidator.Validator
}
