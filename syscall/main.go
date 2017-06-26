package main

import (
	"syscall"
	"log"
	"unsafe"
)

func main(){
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil{
		log.Println(err)
		return
	}
	defer dll.Release()

	proc, err := dll.FindProc("GetTickCount")
	if err != nil{
		log.Println(err)
		return
	}

	//var void interface{} = nil
	ptr := unsafe.Pointer(nil)
	pvoid := uintptr(ptr)
	ticks, _, err := proc.Call(pvoid)

	log.Println(ticks)
}
