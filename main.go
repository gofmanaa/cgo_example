package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L${SRCDIR}/lib -ladd
#include "addlib.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)


func main() {
	counter := uint32(0)

	startTime := time.Now().UTC();
	for {
		counter++
		val := C.add(C.uint32_t(counter))
		valStr := fmt.Sprintf("%d", val)

		res := Concat("prefix_", valStr)

		fmt.Printf("res: %s\n", res)

		//concat returns *C.char

		if counter > 1_000_000 {
			break
		}
	}
	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func Concat(msg string, msg2 string) string {
	cMsg1 := C.CString(msg)
	cMsg2 := C.CString(msg2)
	result := C.concat(cMsg1, cMsg2)

	C.free(unsafe.Pointer(cMsg1))
	C.free(unsafe.Pointer(cMsg2))

	res := C.GoString(result)
	C.free(unsafe.Pointer(result))

	return res
} 

func Concat2(msg string, msg2 string) string {
	cMsg1 := C.CString(msg)
	cMsg2 := C.CString(msg2)
	result := C.concat2(cMsg1, cMsg2)

	C.free(unsafe.Pointer(cMsg1))
	C.free(unsafe.Pointer(cMsg2))

	res := C.GoString(result)
	C.free(unsafe.Pointer(result))

	return res
} 