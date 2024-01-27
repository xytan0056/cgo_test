package main

import (
	"fmt"
	"time"
	"unsafe"

)

/*
#include <security/pam_appl.h>
#include <stdlib.h>
#include <add.h>
*/
import "C"


func running() error {
	fmt.Println("running")
	time.Sleep(time.Second * 6)
	return nil
}


func main() {

	fmt.Println("wow")
	// Call the C function from Go
	result := C.add(3, 7)
	fmt.Printf("Result from C: %d\n", result)
	
	fmt.Println(f( nil))

}


func f(pamh *C.pam_handle_t) error {
	envName := C.CString("SSH_AUTH_INFO_0")
	pamEnv := C.pam_getenv(pamh, envName)
	C.free(unsafe.Pointer(envName))

	env := C.GoString(pamEnv)
	if env == "" {
		return fmt.Errorf("no SSH_AUTH_INFO_0")
	}
	return nil
}
