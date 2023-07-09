package main

import (
	"github.com/boltdb/bolt"
	"rest-api.com/utils"
)

const (
	dbName          = "blockchain.db"
	championsBucket = "champion"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		// bolt.Open : 열기는 지정된 경로에서 DB를 작성하고 열고 없으면 생성
		dbPointer, err := bolt.Open(dbName, 000, nil)
		db = dbPointer
		utils.HandleErr(err)

		// Update : 읽기-쓰기 트랜잭션 실행
		err = db.Update(func(tx *bolt.Tx) error {
			// 없으면 버켓 생성
			_, err := tx.CreateBucketIfNotExists([]byte(championsBucket))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)
	}
	return db
}
