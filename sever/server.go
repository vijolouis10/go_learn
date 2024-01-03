package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandle(w http.ResponseWriter,r *http.Request)  {
	if err:=r.ParseForm();err !=nil{
		fmt.Fprint(w,"parse form err %v",err)
		return
	}
	fmt.Fprint(w, "Post request success")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprint(w,"\nName: ",name,"\n")
	fmt.Fprint(w,"Adress: ",address)
}

// func helloHandle(w http.ResponseWriter,r *http.Request)  {
// 	if r.URL.Path !="/hello"{
// 		http.Error(w,"404 not found",http.StatusNotFound)
// 		return
// 	}
// 	if r.Method !="GET"{
// 		http.Error(w,"method is not supported",http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprint(w,"hello vijo louis")
// }


func main()  {
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandle)
	// http.HandleFunc("/hello",helloHandle)

	fmt.Printf("server is started port 8000\n")
	if err:=http.ListenAndServe(":8000",nil); err !=nil{
		log.Fatal(err)
	}
}