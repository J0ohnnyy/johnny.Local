package main
import "fmt"
func main(){
	var Prixunitair,Qunatite,TauxRemise,remise,totalavantremise,totalht,totalTTC,montantTVA float64
	const TauxTVA float64 =0.2
	fmt.Printf("Quel est le prix unitair?")
	fmt.Scanln(&Prixunitair)
	fmt.Printf("Combient!")
	fmt.Scanln(&Qunatite)
	fmt.Printf("Est votre taux de remise?")
	fmt.Scanln(&TauxRemise)
	totalavantremise= Prixunitair*Qunatite
	fmt.Println("Le total hors-taxe avant remise:", totalavantremise)
    remise= totalavantremise * TauxRemise
	fmt.Println("Le total hors-taxe après remise:",remise)
    totalht= totalavantremise - remise
	fmt.Println("le tottal hors-taxe:",totalht)
	montantTVA= totalht * TauxTVA
	fmt.Println("Le montant de la TVA",montantTVA)
	totalTTC=totalht + montantTVA
	fmt.Println("Total TTC:",totalTTC)
}