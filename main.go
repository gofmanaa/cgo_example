package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L${SRCDIR}/lib -ladd
#include "addlib.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	counter := uint32(0)
	for {
		counter++
		// Call the C add function from Go code
		val := C.add(C.uint32_t(counter))
		fmt.Println(val)
	}
}
