package main

// #include <stdlib.h>
// #include "wrapper.c"
import "C"
import "unsafe"

import (
    "fmt"
    "time"
)

func write(filename string, message string) string {
    f := C.CString(filename)
    defer C.free(unsafe.Pointer(f))
    shm_fd := C.my_shm_open(f,C.int(4096))
    for i:=45;i<55;i++{
    x := C.my_shm_ftruncate(shm_fd,C.int(4096))
    if int(x) == -1 {
	return "FAIL ftruncate"
    }

    	C.my_shm_map_write(shm_fd,C.int(4096),C.CString(message+string(i)))
	time.Sleep(1* time.Second)
    }
    return "DONE"
}


func main() {
	fmt.Println(write("OS","Hello World"))
        

}
