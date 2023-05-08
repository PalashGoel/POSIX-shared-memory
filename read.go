package main

// #include <stdlib.h>
// #include "wrapper.c"
import "C"
import "unsafe"
import "fmt"
import "time"
/*
func read(filename string) string {
    f := C.CString(filename)
    defer C.free(unsafe.Pointer(f))
    s := C.my_shm_read(f)
    defer C.free(unsafe.Pointer(s))
    return C.GoString(s)
}

func readData() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(read("OS"))
	}
}

func main() {
        timer1 := time.NewTimer(60 * time.Second)
        go readData()
        <-timer1.C
}*/


func read(filename string) string {
    f := C.CString(filename)
    defer C.free(unsafe.Pointer(f))
    shm_fd := C.my_shm_open(f,C.int(4096))
    for i:=0;i<10 ;i++{
    	s:= C.my_shm_map_read(shm_fd,C.int(4096))
	fmt.Println(C.GoString(s))
	time.Sleep(1 * time.Second)
    }
    defer C.my_shm_unlink(f)
    return C.GoString(f)
}

func main(){

	fmt.Println(read("OS"))
}
