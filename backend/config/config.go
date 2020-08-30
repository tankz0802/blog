package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	//ServerMode 服务器运行模式
	ServerMode string
	//ServerPort 服务器使用端口
	ServerPort string
	//DB 使用的数据库
	DB         string
	//DBHost 数据库地址
	DBHost     string
	//DBPort 数据库端口
	DBPort     string
	//DBUser 数据库连接用户名
	DBUser     string
	//DBPassword 数据库连接密码
	DBPassword string
	//DBName 数据库名
	DBName     string
	//LogFilePath 日志输出目录
	LogFilePath string
	//LogFilePath 日志文件名
	LogFileName string
	//token过期时间
	//TokenExpiredTime int64
	//token签名的密钥
	TokenSignKey string
	//token签发人
	TokenIssuer string
)

//InitConfig 加载config.ini
func InitConfig(){
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("config.ini加载失败: "+err.Error())
	}
	loadServer(cfg)
	loadMysql(cfg)
	loadLog(cfg)
	loadToken(cfg)
	log.Println("加载config.ini完毕")
}

func loadServer(cfg *ini.File) {
	ServerMode = cfg.Section("server").Key("server_mode").MustString("debug")
	ServerPort = cfg.Section("server").Key("server_port").MustString(":12131")
}

func loadMysql(cfg *ini.File) {
	DB = cfg.Section("mysql").Key("db").MustString("mysql")
	DBHost = cfg.Section("mysql").Key("db_host").MustString("127.0.0.1")
	DBPort = cfg.Section("mysql").Key("db_port").MustString("3306")
	DBUser = cfg.Section("mysql").Key("db_user").MustString("root")
	DBPassword = cfg.Section("mysql").Key("db_password").MustString("123456")
	DBName = cfg.Section("mysql").Key("db_name").MustString("blog")
}

func loadLog(cfg *ini.File) {
	LogFilePath = cfg.Section("log").Key("log_file_path").MustString("log/")
	LogFileName = cfg.Section("log").Key("log_file_name").MustString("system.log")
}

func loadToken(cfg *ini.File) {
	//TokenExpiredTime = cfg.Section("token").Key("token_expired_time").MustInt64(12)
	TokenSignKey = cfg.Section("token").Key("token_sign_key").MustString("blog_token_sign_key")
	TokenIssuer = cfg.Section("token").Key("token_issuer").MustString("tankz")
}