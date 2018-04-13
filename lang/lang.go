package lang

import (
	"runtime"
	"path/filepath"
	"github.com/zpatrick/go-config"
	"os/exec"
	"os"
)

var (
	_lang    *Lang
	langMap map[string]*config.Config
	langFileName = "lang"
	wd string
)

type Lang struct{}

func GetInstance() *Lang {
	if _lang == nil {
		_lang = &Lang{}

	}
	return _lang
}

func loadLangConf(currLang string) *config.Config {
	_, ok := langMap[currLang]
	if ok != true {
		if runtime.GOOS == "windows" {
			wd = getWinConfiPath()
		} else {
			wd = getLinuxConfiPath()
		}
		filepath := wd + "/" + currLang + ".ini"
		iniFile := config.NewINIFile(filepath)
		return config.NewConfig([]config.Provider{iniFile})
	}
	return langMap[currLang]
}

func getWinConfiPath() string {
	_, currFilePath, _, _ := runtime.Caller(0)
	return filepath.Dir(currFilePath)
}

func getLinuxConfiPath() string {
	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	ApplicationDir += langFileName
	return ApplicationDir
}

func (l *Lang) GetTip(currLang string, moduleKey string, strKey string) string {
	hander := loadLangConf(currLang)
	_keyPins := moduleKey + "." + strKey
	currTip, _ := hander.String(_keyPins)
	return currTip
}
