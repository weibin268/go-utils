package io

import "testing"

func TestReadAllText(t *testing.T) {
	text := ReadAllText("/Users/zhuang/Documents/temp/test.txta")
	t.Log(text)
}
