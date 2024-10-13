package database

import (
	constants "amartha-billing-app/common"
	"amartha-billing-app/config"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.GetDbConfig().(*config.DatabaseConfig)
		db, err := gorm.Open(constants.MYSQL_DIALECT, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseConfig.MySQL.DbUsername,
			databaseConfig.MySQL.DbPassword,
			databaseConfig.MySQL.DbHost,
			databaseConfig.MySQL.DbPort,
			databaseConfig.MySQL.DbDatabase,
		))

		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.DB().SetConnMaxIdleTime(5 * time.Minute)
		db.DB().SetConnMaxLifetime(15 * time.Minute)

		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}
		instance = db
	})
	return instance
}
