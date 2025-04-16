package main

import (
	"fmt"
	"testing"
)
/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L${SRCDIR}/lib -ladd
#include "addlib.h"
#include <stdlib.h>
*/

func BenchmarkConcat(b *testing.B) {
	// Setup input for benchmarking
	cMsg1 := "prefix_"

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Use a different value for each iteration
		cMsg2 := fmt.Sprintf("%d", i)

		// Call the concat function and avoid optimizing away
		_ = Concat(cMsg1, cMsg2)
	}
}

func BenchmarkConcat2(b *testing.B) {
	// Setup input for benchmarking
	cMsg1 := "prefix_"

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Use a different value for each iteration
		cMsg2 := fmt.Sprintf("%d", i)

		// Call the concat function and avoid optimizing away
		_ = Concat2(cMsg1, cMsg2)
	}
}