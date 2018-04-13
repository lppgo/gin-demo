package validator

import (
	"demo/modules/common/base"
)

func ValiFactory(valiName string) *base.Validator {
	factory := &base.Validator{}
	switch valiName {
	case "MobileValidator":
		factory = NewMobileValidator()
		break
	case "UserValidator":
		factory = NewUserValidator()
		break
	}
	return factory
}
