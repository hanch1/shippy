package main

// 使用mysql数据库存储user
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shippy/user/model"
	"time"
)

func ConnectMysql(dsn string) (*gorm.DB, error) {

	dsn = "root:root@tcp(127.0.0.1:3306)/shippy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 数据库连接池
	// GORM 使用 database/sql 维护连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})

	return db, err

}
