package main

import "log"

type comal int

func main() {
	a := []int{1, 2}
	b := comal(a[1])
	res, ok := interface{}(b).(comal)
	log.Printf("ok is %v and res %v \n", ok, res)
	switch t := interface{}(b).(type) {
	//switch t := b.(type) {
	case comal:
		log.Printf("its %v\n", t)
	case int:
		log.Printf("its int %v\n", t)
	default:
		log.Printf("out of inspect!")
	}
	if ok == false {
		log.Fatalf("failed, b's type is still string %v \n", b)
	}

	log.Printf("a[1] is %v \n", res)
	log.Printf("a[0] is %v \n", a[0])

}
