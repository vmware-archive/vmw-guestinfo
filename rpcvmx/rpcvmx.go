package rpcvmx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sigma/vmw-guestinfo/rpcout"
)

func ConfigGetString(key string, default_value string) string {
	out, _ := rpcout.SendOne("info-get guestinfo.%s", key)
	if len(out) == 0 {
		return default_value
	}
	return string(out)
}

func ConfigGetBool(key string, default_value bool) bool {
	val := strings.ToLower(ConfigGetString(
		key, fmt.Sprintf("%t", default_value)))
	if val == "true" {
		return true
	} else if val == "false" {
		return false
	} else {
		return default_value
	}
}

func ConfigGetInt(key string, default_value int) int {
	val := ConfigGetString(key, "")

	if val != "" {
		res, err := strconv.Atoi(val)
		if err != nil {
			return default_value
		}
		return res
	} else {
		return default_value
	}
}
