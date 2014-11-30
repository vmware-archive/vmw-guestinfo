package main

import (
	"fmt"

	"github.com/sigma/vmw-guestinfo/rpcvmx"
	"github.com/sigma/vmw-guestinfo/vmcheck"
)

func main() {
	if !vmcheck.IsVirtualWorld() {
		return
	}
	s1 := rpcvmx.ConfigGetString("foo", "foo")
	s2 := rpcvmx.ConfigGetString("bar", "foo")
	fmt.Println(s1)
	fmt.Println(s2)
}
