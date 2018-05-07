package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

//HelperService ...
type HelperService struct{}

// IntToHex Convert integer to hexadecimal
func (h *HelperService) IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
