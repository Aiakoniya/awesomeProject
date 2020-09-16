package main

import "fmt"

type kazakh struct{
}
func(l *kazakh) speak(c *cache){
	fmt.Println("I can speak in Kazakh!")
}
