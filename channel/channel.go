package main

import "fmt"

func print(c chan string)  {
	c<-"vijo"
}

func main() {
	name:=make(chan string)
	go print(name)
	first_name:=<-name
	fmt.Println(first_name)
}
