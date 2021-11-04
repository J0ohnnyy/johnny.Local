package main

import "fmt"
import "math"
func main(){
	var a,b,c,x1,x2,delta float64
	fmt.Println("resolution de l'equation du 2éme degré")
	fmt.Printf("donnez a :")
	fmt.Scanln(&a)
	fmt.Printf("donnez b :")
	fmt.Scanln(&b)
	fmt.Printf("donnez c :")
	fmt.Scanln(&c)
	if a==0{
		if b==0{
			if c==0{
				fmt.Println("Il y l' infinite de solution.")
			} else {
				fmt.Println(" Il n'y a pas de solution.")}
        }else {
				fmt.Println("Il y a une solution :", -c/b)
			}
		} 
    }else {
		delta = b*b-4*a*c
		switch {
		case delta ==0:
			fmt.Println("Il y a une seul solution",-b/(2*a) )
		case delta >0 :
			fmt.Println("Il existe deux solution ")
			fmt.Println("premier solution:",(-b-math.Sqrt(delta))/(2*a) )
			fmt.Println("deuxieme solution:",(-b+math.Sqrt(delta))/(2*a) )
		default:
			fmt.Println("Il n'existe aucune solution")
		}
	}
}