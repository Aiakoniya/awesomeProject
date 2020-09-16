package main

import "fmt"

type english struct{
}
func(l *english) speak(c *cache) {
	fmt.Println("I can speak in English!")
}
