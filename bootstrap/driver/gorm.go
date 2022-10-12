package driver

import (
	"account_check/app/mylogger"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var gErr error

func InitGorm() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s&parseTime=true",
		GVA_VP.GetString("mysql.username"),
		GVA_VP.GetString("mysql.password"),
		GVA_VP.GetString("mysql.host"),
		GVA_VP.GetString("mysql.port"),
		GVA_VP.GetString("mysql.database"),
		GVA_VP.GetString("mysql.charset"),
		GVA_VP.GetString("mysql.timeout"),
	)

	sqlDB, sErr := sql.Open("mysql", dbDSN)

	if sErr != nil {
		log.Println("GORM现有数据库连接失败，GORM功能将不可用。。。", sErr)
		//os.Exit(200)
	} else {
		log.Println("尝试连接GORM... ")
	}

	//记录sql日志
	newLogger := logger.New(
		log.New(mylogger.Logger, "", log.LstdFlags),

		logger.Config{
			SlowThreshold: 1 * time.Second,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)


	GVA_DB, gErr = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
		Logger:                 newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,

		},
	})

	if gErr != nil {
		log.Println("GORM数据库连接失败。。。", gErr)
		//os.Exit(200)
	} else {
		log.Println("GORM已连接现有数据库驱动 >>> ")
	}

}
