package io

import "testing"

func TestJoinPath(t *testing.T) {
	path := JoinPath("/Users/zhuang/Documents/", "temp", "test.txt")
	t.Log(path)
}
