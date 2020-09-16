package main

import "fmt"

type russian struct{
}
func (l *russian) speak(c *cache){
	fmt.Println("I can speak in Russian!")
}
