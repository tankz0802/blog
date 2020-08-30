package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"blog/config"
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	db *gorm.DB
	err error
)

func InitPostgres() {
	//args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err = gorm.Open(config.DB, args)
	if err != nil {
		db.Close()
		panic(err)
	}

	//设置数据表名称为单数
	db.SingularTable(true)

	db.AutoMigrate(&User{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10*time.Second)

	log.Println("postgres connect successfully")
}
