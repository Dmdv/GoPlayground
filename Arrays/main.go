package main

import "log"

func main() {
	log.Println("This is a test")
	bits1 := []int {1,2,3,}
	bits2 := []int {4,5,6,}

	log.Println("Exampe of copy:")
	var bits3 []int = []int{3,3,3,3,}
	copy(bits3, bits2)
	log.Println(bits3)

	log.Println("Exampe of append:")
	bits2 = append(bits2, bits1...)
	log.Println(bits1, bits2)
}
