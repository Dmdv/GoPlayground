package main

import (
	"syscall"
	"unsafe"
	"fmt"
)

// https://stackoverflow.com/questions/33314958/can-a-dll-made-in-c-sharp-be-used-in-a-golang-application
// https://github.com/matiasinsaurralde/go-dotnet

import (
"fmt"
"syscall"
"unsafe"
)

func main() {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("MessageBoxW")
	var MB_YESNOCANCEL = 0x00000003

	ret, _, _ := proc.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
		uintptr(MB_YESNOCANCEL))

	fmt.Printf("Return: %d\n", ret)

}