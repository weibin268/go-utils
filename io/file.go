package io

import (
	"io"
	"log"
	"os"
)

// ReadAllText 读取文本文件
func ReadAllText(filePath string) string {
	var text string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	text = string(data)
	return text
}
