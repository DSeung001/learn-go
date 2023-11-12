package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0 ")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0 ")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1 ")
	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 3 ")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233 ")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0 ")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0 ")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1 ")
	assert.Equal(2, fibonacci2(3), "fibonacci2(3) should be 3 ")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233 ")
}

// 성능 검사
func BenchmarkFibonacci1(b *testing.B) {
	// N 값을 go에서 적당히 늘려가면서 성능 테스트
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

// 재귀보다 반복문이 훨씬 빠른걸 알 수 있음
func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
