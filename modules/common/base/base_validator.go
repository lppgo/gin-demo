package base

type Validator struct {
	Sence     map[string][]string
	Rules     map[string][]string
	Messages  map[string][]string
	currSence string
}

func (validator *Validator) SetSence(sence string) {
	validator.currSence = sence
}

func (validator *Validator) GetCurrSence() string {
	return validator.currSence
}

func (validator *Validator) GetSenceFields(sence string) []string {
	return validator.Sence[sence]
}

func (validator *Validator) GetRules() map[string][]string {
	tempData := map[string][]string{}
	tempData = validator.Rules
	if validator.currSence != "" {
		tempData = validator.MatchSence(validator.Rules)
	}
	return tempData
}

func (validator *Validator) GetMessages() map[string][]string {
	tempData := map[string][]string{}
	tempData = validator.Messages
	if validator.currSence != "" {
		tempData = validator.MatchSence(validator.Messages)
	}
	return tempData
}

func (validator *Validator) MatchSence(datas map[string][]string) map[string][]string {
	tempData := map[string][]string{}
	senceLists, ok := validator.Sence[validator.GetCurrSence()]
	if ok == false {
		return datas
	}
	for _, row := range senceLists {
		currRow, ok := datas[row]
		if ok == true {
			tempData[row] = currRow
		}
	}
	return tempData
}
