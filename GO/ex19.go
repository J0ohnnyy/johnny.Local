package main

import "fmt"

func main(){
  var n,F int
   fmt.Printf("tapez un nombre:")
   fmt.Scanln(&n)
   F=1
   for i:=1;i<=n;i++{
	   F=i*F
   }
   fmt.Printf("Le factorielle:",F)
}