// package main

// import (
// 	"fmt"
// )

// func main()  {
// 	fmt.Printf("%.0f\n",30000.300003)
// }

package main  
import (  
   "fmt"  
)  
func main() {  
   var input int  
Loop:   
   fmt.Print("Enter your age: ")  
   fmt.Scanln(&input)  
   if (input <= 17) { 
	  fmt.Println("You are not eligible to vote ")  
      goto Loop  
   } else {  
      fmt.Print("You can vote ")  
   }  
}  
