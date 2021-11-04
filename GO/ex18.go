package main 

import "fmt"

func main(){

var n int
 
fmt.Printf("tapez un nombre:")
fmt.Scanln(&n)
 
   for i:=n+1;i<=n+10;i++{
	fmt.Println(i)
    }
}