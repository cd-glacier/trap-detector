package main

import "log"

func main() {
	s := []string{"hoge", "foo"}
	log.Printf("before append: %p", s)
	expandSlice(s)

	log.Printf("%#v", s)
}

func expandSlice(s []string) {
	log.Printf("before append: %p", s)
	s = append(s, "expand")
	log.Printf("after append: %p", s)
}
