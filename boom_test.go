package cgo_static_boom

import (
	"net"
	"os/user"
	"testing"
)

func TestBoom(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:5423")
	user.Current()
	t.Fatalf("conn: %s, err: %s", conn, err)
}
