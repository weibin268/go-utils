package io

import "testing"

func TestReadText(t *testing.T) {
	text := ReadText("/Users/zhuang/Documents/temp/test.txt")
	t.Log(text)
}

func TestWriteText(t *testing.T) {
	WriteText("/Users/zhuang/Documents/temp/test.txt", "aaaa")
}
