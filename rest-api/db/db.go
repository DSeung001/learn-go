package db

import (
	"github.com/boltdb/bolt"
	"rest-api.com/utils"
)

// dbName : DB 명으로 해당 이름으로 파일이 생성됩니다
// championBucket : Bucket 은 RDBMS 의 테이블과 같은 역할을 합니다.
const (
	dbName          = "champions.db"
	championsBucket = "champion"
)

var db *bolt.DB

// DB : db 커넥션
func DB() *bolt.DB {
	if db == nil {
		// bolt.Open : dbName 으로 데이터베이스 열기, 뒤에 0600 은 파일 권한
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)

		// Update : 읽기-쓰기 트랜잭션 실행
		err = db.Update(func(tx *bolt.Tx) error {
			// CreateBucketIfNotExists : 해당 버켓명이 없으면 버켓 생성
			_, err := tx.CreateBucketIfNotExists([]byte(championsBucket))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

// Close : db connection 종료
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
		// 커서로 bucket 내부를 돌며 map 에 저장
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
		// id가 존재하는 값이면 값 덮어씌우기
		err := bucket.Put(utils.IntToBytes(id), []byte(name))
		return err
	}))
}

// DeleteChampion : DB에 특정 아이디로 데이터 삭제
func DeleteChampion(id int) {
	utils.HandleErr(DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(championsBucket))
		// id로 데이터 삭제
		err := bucket.Delete(utils.IntToBytes(id))
		return err
	}))
}
