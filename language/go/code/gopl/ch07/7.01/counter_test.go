package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	c := new(Counter)
	fmt.Fprintf(c, "Byte Count\n%v", "xbu")
	fmt.Println(*c)
}
