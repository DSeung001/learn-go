package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test -run Square
// Square로 시작하는 함수만 테스트

func TestSquare1(t *testing.T) {
	// Equal 외에도 NotEqual, Greater, Len, Nil, NotNil 등 다양한 함수가 있음
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
