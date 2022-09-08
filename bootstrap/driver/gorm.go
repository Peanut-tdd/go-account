package driver

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var gErr error

func InitGorm() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		GVA_VP.GetString("mysql.username"),
		GVA_VP.GetString("mysql.password"),
		GVA_VP.GetString("mysql.host"),
		GVA_VP.GetString("mysql.port"),
		GVA_VP.GetString("mysql.database"),
		GVA_VP.GetString("mysql.chartset"),
		GVA_VP.GetString("mysql.timeout"),
	)

	sqlDB, sErr := sql.Open("mysql", dbDSN)

	if sErr != nil {
		log.Println("GORM现有数据库连接失败，GORM功能将不可用。。。", sErr)
		//os.Exit(200)
	} else {
		log.Println("尝试连接GORM... ")
	}

	GVA_DB, gErr = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
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
