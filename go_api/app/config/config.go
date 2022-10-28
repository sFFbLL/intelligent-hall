package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// 应用配置项
var cfgApplication *viper.Viper

//载入配置文件
func Setup(path string) {
	// 加载配置文件
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

}