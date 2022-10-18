package driver

// mysql文档：http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/mysql%E4%BD%BF%E7%94%A8.html

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var MysqlDbErr error

func InitMysql() {
	log.Println("尝试连接MySQL服务...")

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		GVA_VP.GetString("mysql.username"),
		GVA_VP.GetString("mysql.password"),
		GVA_VP.GetString("mysql.host"),
		GVA_VP.GetString("mysql.port"),
		GVA_VP.GetString("mysql.database"),
		GVA_VP.GetString("mysql.charset"),
		GVA_VP.GetString("mysql.timeout"),
	)

	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}

	//max open connections
	dbMaxOpenConns, _ := strconv.Atoi(GVA_VP.GetString("mysql.max_open_cons"))
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(GVA_VP.GetString("mysql.max_idle_cons"))
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	// dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
	// MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		log.Println("MySQL数据库连接失败。。。", MysqlDbErr.Error())
		//os.Exit(200)
	} else {
		log.Println("MySQL已连接 >>> ")
	}
}
