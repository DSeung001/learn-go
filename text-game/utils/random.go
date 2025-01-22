package utils

import (
	"math/rand"
)

// 랜덤 범위 계산 함수
func GetRandomValue(max int) int {
	if max <= 0 {
		return 0
	}
	return rand.Intn(max) + 1
}
