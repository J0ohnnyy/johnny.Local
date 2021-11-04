package main 
import "fmt"

func main(){
	var n,S int

	fmt.Printf("tapez un nombres:")
	fmt.Scanln(&n)
	S=0
	for i:=1; i<=n; i++ {
		S=S+i
		
	}
	fmt.Println("votres sommes est:",S)
}