package vmcheck

/*
#cgo CFLAGS: -I../include
#include <stdlib.h>
*/
import "C"

func IsVirtualWorld() bool {
	return true
}
