package config

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"github.com/zpatrick/go-config"
)

var gopath string
var currPath string
var err error
var env string
var port string
var configHandle *config.Config
var envHandle *config.Config

var configFileName = "config"

func init() {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		currPath = getWinConfiPath()
	} else {
		currPath = getLinuxConfiPath()
	}

	env = getEnv()
	getSettings()
	getCantants()

}

func getWinConfiPath() string {
	_, currFilePath, _, _ := runtime.Caller(0)
	return filepath.Dir(currFilePath)
}

func getLinuxConfiPath() string {
	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	ApplicationDir += configFileName
	return ApplicationDir
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func getGopath() string {
	return os.Getenv("GOPATH")
}

// GetENV is
func GetENV() string {
	return env
}

// GetPORT is
func GetPORT() string {
	return port
}

func getEnv() string {
	filepath := currPath + "/env.ini"
	iniFile := config.NewINIFile(filepath)
	c := config.NewConfig([]config.Provider{iniFile})
	env, _ = c.String("APP.ENV")
	port, _ = c.String("APP.PORT")
	return env
}

func GetGconfig(gName string, key string) string {
	filepath := currPath + "/" + env + ".ini"
	iniFile := config.NewINIFile(filepath)
	envHandle = config.NewConfig([]config.Provider{iniFile})
	result, _ := envHandle.String(gName + "." + key)
	return result
}

func getSettings() {
	filepath := currPath + "/" + env + ".ini"
	iniFile := config.NewINIFile(filepath)
	envHandle = config.NewConfig([]config.Provider{iniFile})
}

func getCantants() {
	filepath := currPath + "/config.ini"
	// fmt.Println(filepath)
	iniFile := config.NewINIFile(filepath)
	configHandle = config.NewConfig([]config.Provider{iniFile})

}

// GetConfigs is
func GetConfigs() *config.Config {
	return configHandle
}

// GetENVConfigs is
func GetENVConfigs() *config.Config {
	return envHandle
}
