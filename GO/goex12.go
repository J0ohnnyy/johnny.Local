package main
import "fmt"
func main(){
	var a int
	fmt.Printf("donnez un nombre compris entre 10 et 20:")
	fmt.Scanln(&a)
	switch{
	  case a<=10 && a>=20 :
		fmt.Println("Bon choix!")
	  case a<10 :
		fmt.Println("Plux petie!")
	  case a>20 :
		fmt.Println("Plux grand!")		
	}
}
