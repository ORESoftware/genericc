package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "../genericc/greeter.h"
// #include "../genericc/greeter.c"
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

func greet(name string, year int) string {
  return fmt.Sprintf("Greetings, %s from %d! We come in peace :)", name, year);
}

func main() {

	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))
	year := C.int(2018)

	{
		now := time.Now().UnixNano();

		for i := 1; i <= 100000; i++ {

			greet("Gopher", 2018)

			//b := C.GoBytes(ptr, size)
			//fmt.Println(string(b))
		}

		fmt.Println("1:",((time.Now().UnixNano() - now)/1000000))
	}

	{
		now := time.Now().UnixNano();

		for i := 1; i <= 100000; i++ {

			 C.greet(name, year, (*C.char)(ptr))
			//b := C.GoBytes(ptr, size)
			//fmt.Println(string(b))
		}

		fmt.Println("2:",((time.Now().UnixNano() - now)/1000000))
	}




}