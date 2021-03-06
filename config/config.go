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
	JwtTime   int
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

var ServerSettings = &Server{}
var AppSetting = &App{}
var MysqlSettings = &Mysql{}
var cfg *ini.File

//look into setting them into variables, to use os.Env
func Setup(path string) {
	var err error
	cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("setting: %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSettings)
	mapTo("mysql", MysqlSettings)

	ServerSettings.ReadTimeout = ServerSettings.ReadTimeout * time.Second
	ServerSettings.WriteTimeout = ServerSettings.WriteTimeout * time.Second
}

// load each configurations
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
