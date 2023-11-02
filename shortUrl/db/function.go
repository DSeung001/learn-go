package db

import (
	"database/sql"
	"log"
	"shortUrl.com/model"
	"shortUrl.com/utils"
)

func GetUrlList() []model.Url {
	var urls []model.Url

	rows, err := getDBConnection().Query("SELECT * FROM url")
	utils.HandleErr(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		utils.HandleErr(err)
	}(rows)

	for rows.Next() {
		var url model.Url
		err := rows.Scan(&url.Id, &url.AliasURL, &url.FullURL)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, url)
	}
	return urls
}

func InsertUrl(url model.Url) {
	stmt, err := getDBConnection().Prepare("INSERT INTO url(alias_url, full_url) VALUES(?, ?)")
	utils.HandleErr(err)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		utils.HandleErr(err)
	}(stmt)

	_, err = stmt.Exec(url.AliasURL, url.FullURL)
	utils.HandleErr(err)
}

func PatchUrl(url model.Url, id string) {
	stmt, err := getDBConnection().Prepare("UPDATE url SET alias_url=?, full_url=? WHERE id=?")
	utils.HandleErr(err)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		utils.HandleErr(err)
	}(stmt)

	_, err = stmt.Exec(url.AliasURL, url.FullURL, id)
	utils.HandleErr(err)
}

func DeleteUrl(id string) {
	stmt, err := getDBConnection().Prepare("DELETE FROM url WHERE id=?")
	utils.HandleErr(err)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		utils.HandleErr(err)
	}(stmt)

	_, err = stmt.Exec(id)
	utils.HandleErr(err)
}
