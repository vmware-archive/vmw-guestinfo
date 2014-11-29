package main

import (
	"fmt"

	"github.com/sigma/vmw-guestinfo/rpcvmx"
)

func main() {
	s1 := rpcvmx.ConfigGetString("foo", "foo")
	s2 := rpcvmx.ConfigGetString("bar", "foo")
	fmt.Println(s1)
	fmt.Println(s2)
}
