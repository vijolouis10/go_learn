package main

import (
	"fmt"
	"os"
)

func main()  {
	file,err:=os.Open("./text.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("sucess",file.Name())
}