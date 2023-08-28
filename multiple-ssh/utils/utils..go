package utils

import (
	"log"
)

// HandleErr : 에러 발생시 Panic 으로 Log 출력
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
