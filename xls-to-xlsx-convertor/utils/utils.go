package utils

import "log"

var logFn = log.Panic

// HandleErr : 간단하게 에러 처리
func HandleErr(err error) {
	if err != nil {
		logFn(err)
	}
}
