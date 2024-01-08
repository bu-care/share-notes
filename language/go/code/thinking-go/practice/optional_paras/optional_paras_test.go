package optional_paras

import (
	"fmt"
	"testing"
)

func TestOptionParas(t *testing.T) {
	srv := NewServer("127.0.0.1", 8000, Protocol("udp"), TimeoutOpt(10))
	fmt.Println(srv)
}
