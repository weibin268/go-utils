package cmd

import (
	"testing"
)

func TestCall(t *testing.T) {
	str, _ := Call("cmd", "/C", "dir")
	t.Log(str)
}
