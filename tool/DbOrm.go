package tool

import (
	"fmt"
	"gogochat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Orm struct {
	*gorm.DB
}

var DBEngine = new(Orm)

func InitDB(config *Config) *Orm {
	//sql打印日志
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      logger.Info,
	})

	dbConfig := config.Database
	connstr := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("----------connstr", connstr)
	db, err := gorm.Open(mysql.Open(connstr), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	newOrm := Orm{}
	newOrm.DB = db

	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		return nil
	}
	DBEngine.DB = db

	return &newOrm

}
