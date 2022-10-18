package bytes

import (
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {
	bytes := ToBytes(123)
	fmt.Printf("%v", bytes)
}

func TestToInt(t *testing.T) {
	b := ToBytes(123)
	i := ToInt(b)
	fmt.Printf("%v", i)
}
