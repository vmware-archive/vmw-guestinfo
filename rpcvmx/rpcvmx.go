package rpcvmx

import (
	"github.com/sigma/vmw-guestinfo/rpcout"
)

func ConfigGetString(key string, default_value string) string {
	out, _ := rpcout.SendOne("info-get guestinfo.%s", key)
	if out == "" {
		return default_value
	}
	return out
}
