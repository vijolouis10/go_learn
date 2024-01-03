package main

import (
	"fmt"
	"os"
)

func main()  {
	status:="error"
	if status=="error"{
		fmt.Println("error")
		os.Exit(404)
	}
}