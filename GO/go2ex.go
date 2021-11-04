package main

import "fmt"

 func main() {
	 var a, b, c int
	 fmt.Printf("donnez le debut d'interval:")
	 fmt.Scanln(&a)
	 fmt.Printf("donnez la fin d'interval:")
	 fmt.Scanln(&b)
	 fmt.Printf("tapez la valeur d'un nombre:")
	 fmt.Scanln(&c)
	 if a<c && c<b{
		 fmt.Printf("oui")
	 } else {
		fmt.Printf("non")
	 }

 }