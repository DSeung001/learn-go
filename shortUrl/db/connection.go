package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"shortUrl.com/utils"
)

type Url struct {
	Id       int
	AliasURL string
	FullURL  string
}

func init() {
	utils.HandleErr(getDBConnection().Ping())
	fmt.Println("DB Connected!")
	fmt.Println("DB")
}

func GetUrlList() []Url {
	var urls []Url

	rows, err := getDBConnection().Query("SELECT * FROM url")
	utils.HandleErr(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		utils.HandleErr(err)
	}(rows)

	for rows.Next() {
		var url Url
		err := rows.Scan(&url.Id, &url.AliasURL, &url.FullURL)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, url)
	}
	fmt.Println(urls)
	return urls
}

func getDBConnection() *sql.DB {
	// sql.Open의 경우 query 실행시 DB에 연결됨
	db, err := sql.Open("mysql", getDataSourceName())
	utils.HandleErr(err)
	return db
}

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
