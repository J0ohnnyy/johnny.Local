package main

import "fmt"
func main(){
	var a,b int
	fmt.Printf("Donnez la valeur du premier nombre:")
	fmt.Scanln(&a)
	fmt.Printf("Donnez la valeur du deuxiem nombre:")
	fmt.Scanln(&b)
	if (a<0 && b<0) || (a >0 && b>0){
		fmt.Println("le produit des deux nombres est positif")
	} else{fmt.Println("le produit des deux nombres est negative")
	}
}