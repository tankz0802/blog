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
)

//InitConfig 加载config.ini
func InitConfig(){
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal("config.ini加载失败: "+err.Error())
	}
	loadServer(cfg)
	loadPostgres(cfg)

	log.Println("加载config.ini完毕")
}

func loadServer(cfg *ini.File) {
	ServerMode = cfg.Section("server").Key("server_mode").MustString("debug")
	ServerPort = cfg.Section("server").Key("server_port").MustString(":12131")
}

func loadPostgres(cfg *ini.File) {
	DB = cfg.Section("postgres").Key("db").MustString("postgres")
	DBHost = cfg.Section("postgres").Key("db_host").MustString("106.52.211.246")
	DBPort = cfg.Section("postgres").Key("db_port").MustString("5432")
	DBUser = cfg.Section("postgres").Key("db_user").MustString("blog")
	DBPassword = cfg.Section("postgres").Key("db_password").MustString("blog1214")
	DBName = cfg.Section("postgres").Key("db_name").MustString("blog")
}