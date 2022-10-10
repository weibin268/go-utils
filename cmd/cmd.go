package cmd

import (
	"bytes"
	"os/exec"
)

// Call 调用终端命令
func Call(name string, arg ...string) (string, error) {
	var buff bytes.Buffer
	var errBuff bytes.Buffer
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &buff
	cmd.Stderr = &errBuff
	err := cmd.Run()
	if err != nil && errBuff.Len() > 0 {
		return string(errBuff.Bytes()), err
	} else {
		return string(buff.Bytes()), err
	}
}
