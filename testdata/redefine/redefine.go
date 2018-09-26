package main

import (
	"log"
)

func main() {
	localVariable1()
	localVariable2()
}

func localVariable1() {
	v := "variable"

	if v == "variable" {
		v := "changed variable"

		log.Println(v)
	}

	log.Println(v)
}

func localVariable2() {
	if x := 10; x == -1 {
		log.Println("x block")
	} else if y := 20; y == -1 {
		log.Println("y block")
	} else {
		log.Printf("(x, y) = (%d, %d)\n", x, y)
	}
}
