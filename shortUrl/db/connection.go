package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"shortUrl.com/utils"
)

// init : 패키지 로드시 DB 연결 테스트
func init() {
	utils.HandleErr(getDBConnection().Ping())
	fmt.Println("DB Connected!")
}

// getDBConnection : DB connectionOpener 반환
func getDBConnection() *sql.DB {
	// sql.Open의 경우 query 실행시에만 DB에 연결됨
	db, err := sql.Open("mysql", getDataSourceName())
	utils.HandleErr(err)
	return db
}

// getDataSourceName : DB 연결 정보 반환
func getDataSourceName() string {
	utils.HandleErr(godotenv.Load(".env"))
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName: os.Getenv("DBNAME"),
	}
	return cfg.FormatDSN()
}
