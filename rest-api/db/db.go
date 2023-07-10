package db

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
		// bolt.Open : 열기는 지정된 경로에서 DB를 작성하고 열고 없으면 생성, mode 값으로 파일권한 부여
		dbPointer, err := bolt.Open(dbName, 0600, nil)
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

func Close() {
	DB().Close()
}

// SaveChampion : 챔피언을 DB에 추가
func SaveChampion(name string) {
	utils.HandleErr(DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(championsBucket))

		// championBucket 의 다음 값 가져오기
		id, _ := bucket.NextSequence()
		// 저장
		err := bucket.Put(utils.IntToBytes(int(id)), []byte(name))
		return err
	}))
}

// ReadChampions : DB에 전체 챔피언 가져오기
func ReadChampions() map[int]string {
	m := make(map[int]string)
	utils.HandleErr(DB().View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		bucket := tx.Bucket([]byte(championsBucket))
		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			m[utils.BytesToInt(key)] = string(value)
		}
		return nil
	}))
	return m
}

// UpdateChampion : DB에 특정 아이디 이름 업데이트
func UpdateChampion(id int, name string) {
	utils.HandleErr(DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(championsBucket))
		err := bucket.Put(utils.IntToBytes(id), []byte(name))
		return err
	}))
}

// DeleteChampion : DB에 특정 아이디로 데이터 삭제
func DeleteChampion(id int) {
	utils.HandleErr(DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(championsBucket))
		err := bucket.Delete(utils.IntToBytes(id))
		return err
	}))
}
