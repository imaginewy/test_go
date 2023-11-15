package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Sever struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	Cfg             *ini.File
	ServerSetting   = &Sever{}
	AppSetting      = &App{}
	DatabaseSetting = &Database{}
	RedisSetting    = &Redis{}
)

func Setup() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// func LoadBase() {
// 	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
// }
// func LoadServer() {
// 	sec, err := Cfg.NewSection("server")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'server': %v", err)
// 	}
// 	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
// 	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
// 	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
// }

// func LoadApp() {
// 	sec, err := Cfg.NewSection("app")
// 	if err != nil {
// 		log.Fatalf("Fail to get section 'server': %v", err)
// 	}
// 	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
// 	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
// }
