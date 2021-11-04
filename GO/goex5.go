package main
import "fmt"
func main(){
	var a,b,c int 
	fmt.Printf(" donnez la valeur du a:")
	fmt.Scanln(&a)
	fmt.Printf(" donnez la valeur du b:")
	fmt.Scanln(&b)
	fmt.Printf(" donnez la valeur du c:")
	fmt.Scanln(&c)
	if (a<= b) && (b<=c){
		fmt.Println("Les nombres sont ordonnés.")
	}else {fmt.Println("Les nombres ne sont pas ordonnés.")
	}
}