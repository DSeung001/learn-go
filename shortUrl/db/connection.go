package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"shortUrl.com/utils"
)

// 전역변수
var (
	ctx context.Context
	db  *sql.DB
)

type Url struct {
	Id       int
	AliasURL string
	FullURL  string
}

func init() {
	fmt.Println("init")
	utils.HandleErr(godotenv.Load(".env"))

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT"),
		DBName: os.Getenv("DBNAME"),
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	utils.HandleErr(db.Ping())
	fmt.Println("DB Connected!")
}

func GetUrlList() []Url {
	var urls []Url
	rows, err := db.Query("SELECT * FROM url")
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
	return urls
}
