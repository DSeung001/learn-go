package db

import (
	"database/sql"
	"log"
	"shortUrl.com/model"
	"shortUrl.com/utils"
)

// InsertUrl : url 정보를 DB에 저장
func InsertUrl(url model.Url) {
	stmt, err := getDBConnection().Prepare("INSERT INTO url(alias_url, full_url) VALUES(?, ?)")
	utils.HandleErr(err)

	// defer : 함수가 종료되기 직전에 stmt 종료
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		utils.HandleErr(err)
	}(stmt)

	// Exec : 쿼리 실행
	_, err = stmt.Exec(url.AliasURL, url.FullURL)
	utils.HandleErr(err)
}

// GetUrlList : DB에 저장된 url 정보를 []model.url로 반환
func GetUrlList() []model.Url {
	var urls []model.Url

	rows, err := getDBConnection().Query("SELECT * FROM url")
	utils.HandleErr(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		utils.HandleErr(err)
	}(rows)

	// row마다 돌며 url 정보를 urls에 저장
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

// PatchUrl : DB에 저장된 url 정보를 업데이트
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

// DeleteUrl : DB에 저장된 url 정보를 삭제
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
