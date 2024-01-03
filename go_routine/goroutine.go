package main

import (
	"fmt"
	"time"
)



func main()  {
	go foo("1st goroutine")
	go foo("2nd gorountine")

	time.Sleep(time.Second)
	fmt.Println("FINESHED")
}

func foo(s string)  {
	for i:=0;i<=3;i++{
		time.Sleep(100*time.Microsecond)
		fmt.Println(s,":",i)
	}
}

