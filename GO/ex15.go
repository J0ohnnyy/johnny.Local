package main

import (
	"fmt"
)

func main(){
	var nbH float64 
	var Total, montant, arendr, nb10, nb5, nb2, nb1, Reste int
	const pH float64 = 2 
	fmt.Printf("Entrez le nombre d'heure:")
	fmt.Scanln(&nbH)
	Total= int(nbH*pH)
	fmt.Println("votre somme est:",Total)
	fmt.Printf("Entrez votre montant:")
	fmt.Scanln(&montant)
	arendr =montant-Total
	nb10 = arendr / 10
	Reste = arendr % 10
	nb5=Reste / 5
	Reste=Reste % 5
	nb2=Reste / 2
	Reste=Reste % 2
	fmt.Println("voi-ci votre rest :",nb10,nb5,nb2,nb1)
}