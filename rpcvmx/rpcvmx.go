package rpcvmx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sigma/vmw-guestinfo/rpcout"
)

func ConfigGetString(key string, default_value string) (string, error) {
	out, ok, err := rpcout.SendOne("info-get guestinfo.%s", key)
	if err != nil {
		return "", err
	} else if !ok {
		return default_value, nil
	}
	return string(out), nil
}

func ConfigGetBool(key string, default_value bool) (bool, error) {
	val, err := ConfigGetString(key, fmt.Sprintf("%t", default_value))
	if err != nil {
		return false, err
	}
	switch strings.ToLower(val) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return default_value, nil
	}
}

func ConfigGetInt(key string, default_value int) (int, error) {
	val, err := ConfigGetString(key, "")
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return default_value, nil
	}
	return res, nil
}
