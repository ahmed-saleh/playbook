package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

/*
	App configurations
*/
type App struct {
	JwtSecret string
	// PageSize  int
	PrefixUrl       string
	RuntimeRootPath string

	// ImageSavePath  string
	// ImageMaxSize   int
	// ImageAllowExts []string

	// ExportSavePath string
	// QrCodeSavePath string
	// FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	// TimeFormat  string
}

/*
	Server configurations
*/

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

/*
	Mysql DB configurations
*/
type Mysql struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

var ServerSetting = &Server{}
var AppSetting = &App{}
var MysqlSettings = &Mysql{}
var cfg *ini.File

func Setup(path string) {
	var err error
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("setting: %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("mysql", MysqlSettings)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// load each configurations
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
