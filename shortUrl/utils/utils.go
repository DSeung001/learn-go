// Package utils contains functions to be used across the application
package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var logFn = log.Panic

// HandleErr : 에러가 발생시 처리
func HandleErr(err error) {
	if err != nil {
		fmt.Println("err", err)
		logFn(err)
	}
}

// GetThirdIndexUrl : url의 세번째 인덱스 반환
func GetThirdIndexUrl(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return ""
	} else {
		return parts[2]
	}
}
