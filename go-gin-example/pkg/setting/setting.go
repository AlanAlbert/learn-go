package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	AppMode string

	HttpPort int

	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init() {

	fmt.Println("Module \"setting\" start init...")

	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse \"conf/app.ini\": %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()

	fmt.Println("Module \"setting\" init complete.")
}

func LoadBase() {
	AppMode = Cfg.Section("").Key("APP_MODE").MustString("debug")
}

func LoadServer() {
	section, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section \"server\": %v", err)
	}

	HttpPort = section.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	section, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section \"app\": %v", err)
	}

	PageSize = section.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = section.Key("JWT_SECRET").MustString("")
}
