package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func IntToBytes(number int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, number)
	HandleErr(err)
	return buf.Bytes()
}
