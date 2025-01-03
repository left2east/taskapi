package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var postgresdb *gorm.DB

var err error

func InitPostgresDb() {
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"mygin",
		"8WhJCBXETvrC",
		"template1",
	)

	postgresdb, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
		//关闭自动生成复数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("连接数据库失败，请检查参数：", err)
		panic(err)
	}
	postgresdb.AutoMigrate(&TaskTable{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := postgresdb.DB()
	if err != nil {
		fmt.Println("获取通用数据库对象，使用连接池失败，请检查错误：", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(59 * time.Second)
}
