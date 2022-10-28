package config

import "github.com/spf13/viper"

type Application struct {
	ReadTimeout         int
	WriterTimeout       int
	Host                string
	Port                string
	Name                string
	Mode                string
	StaticFileUrl       string
	StaticPath          string
	Version             string
	TerminalTcp         string
	TerminalHttp        string
	TerminalHttpLogin   string
	TerminalHttpList    string
	TerminalQueryList   string
	TerminalOperateList string
	EnableDP            bool
}

// InitApplication 初始化app配置
func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:         cfg.GetInt("readTimeout"),
		WriterTimeout:       cfg.GetInt("writerTimeout"),
		Host:                cfg.GetString("host"),
		Port:                cfg.GetString("port"),
		Name:                cfg.GetString("name"),
		Mode:                cfg.GetString("mode"),
		StaticFileUrl:       cfg.GetString("staticfileurl"),
		StaticPath:          cfg.GetString("staticpath"),
		Version:             cfg.GetString("version"),
		TerminalTcp:         cfg.GetString("terminalTcp"),
		TerminalHttp:        cfg.GetString("terminalHttp"),
		TerminalHttpLogin:   cfg.GetString("terminalHttpLogin"),
		TerminalHttpList:    cfg.GetString("terminalHttpList"),
		TerminalQueryList:   cfg.GetString("terminalQueryList"),
		TerminalOperateList: cfg.GetString("terminalOperateList"),
		EnableDP:            cfg.GetBool("enabledp"),
	}
}

var ApplicationConfig = new(Application)
