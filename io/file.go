package io

import (
	"io"
	"log"
	"os"
)

// ReadText 读取文本文件
func ReadText(filePath string) string {
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

func WriteText(filePath string, text string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(file, text)
	if err != nil {
		log.Fatal(err)
	}
}
