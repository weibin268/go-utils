package io

import (
	"path/filepath"
)

// JoinPath 路径拼接
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
