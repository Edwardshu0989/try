package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n\nkkkkk")
	err := ioutil.WriteFile("test.txt", d1, 0644)
	check(err)
}
