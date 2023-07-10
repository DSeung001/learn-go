package utils

import (
	"encoding/binary"
	"log"
)

// HandleErr : 에러 발생시 Panic 으로 Log 출력
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// IntToBytes : int 타입을 []byte 타입으로 변경 후 반환
func IntToBytes(integer int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(integer))
	return bytes
}

// BytesToInt : []byte 타입을 int 타입으로 변경 후 반환
func BytesToInt(bytes []byte) int {
	return int(binary.BigEndian.Uint64(bytes))
}
