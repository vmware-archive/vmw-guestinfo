package main

import (
	"fmt"

	"github.com/sigma/vmw-guestinfo/rpcvmx"
	"github.com/sigma/vmw-guestinfo/vmcheck"
)

func main() {
	if !vmcheck.IsVirtualWorld() {
		fmt.Println("not in a virtual world... :(")
		return
	}

	version, typ := vmcheck.GetVersion()
	fmt.Println(version, typ)

	fmt.Println(rpcvmx.ConfigGetString("foo", "foo"))
	fmt.Println(rpcvmx.ConfigGetString("bar", "foo"))
}
