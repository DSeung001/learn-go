// Package utils contains functions to be used across the application
package utils

import (
	"fmt"
	"log"
)

var logFn = log.Panic

// HandleErr : 에러가 발생시 처리
func HandleErr(err error) {
	if err != nil {
		fmt.Println("err", err)
		logFn(err)
	}
}
