package bytes

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ToBytes(n int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, int32(n))
	if err != nil {
		log.Fatal(err)
	}
	return bytesBuffer.Bytes()
}

func ToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		log.Fatal(err)
	}
	return int(x)
}
