package validator

import "demo/modules/common/base"

type UserValidator struct {
	base.Validator
}

func NewUserValidator() *base.Validator {
	userValidator := &UserValidator{}
	userValidator.Sence = map[string][]string{
		"login": []string{"phone", "password"},
	}

	userValidator.Messages = map[string][]string{
		"phone":    []string{"required:uRequired"},
		"password": []string{"required:passRequired", "between:betweenPass"},
	}

	userValidator.Rules = map[string][]string{
		"phone":    []string{"required"},
		"password": []string{"required", "between:6,20"},
	}
	return &userValidator.Validator
}
