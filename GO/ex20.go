package main 

import "fmt"

func main(){
	var n,max,nbr,pos int
	fmt.Printf("Donnez un nombre:")
	fmt.Scanln(&nbr)
	max=nbr
	n=1
	pos=1
	for i:=1;i<=n;i++{
		fmt.Printf("Donnez un nombre:")
		fmt.Scanln(&nbr)
		if nbr>max {
			max=nbr
			pos=i
		}
	}
    fmt.Println("Le nombre le plus grand est :",max)
	fmt.Println("La position du plus grand nombre est:",pos)	
}