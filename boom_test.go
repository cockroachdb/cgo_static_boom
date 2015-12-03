package cgo_static_boom

import (
	"net"
	"os/user"
	"testing"
)

func TestBoom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, _ = net.Dial("tcp", "localhost:1337")
		_, _ = user.Current()
	}
}
