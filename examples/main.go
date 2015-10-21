package main

import (
	"fmt"

	"github.com/vmware/vmw-guestinfo/rpcvmx"
	"github.com/vmware/vmw-guestinfo/vmcheck"
)

func main() {
	if !vmcheck.IsVirtualWorld() {
		fmt.Println("not in a virtual world... :(")
		return
	}

	version, typ := vmcheck.GetVersion()
	fmt.Println(version, typ)

	config := rpcvmx.NewConfig()
	fmt.Println(config.GetString("foo", "foo"))
	fmt.Println(config.GetString("bar", "foo"))
}
