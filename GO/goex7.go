package main
import "fmt"
func main(){
	var a,m int
	fmt.Printf("tapez votre note:")
	fmt.Scanln(&a)
	m = a
	if m>=10{
		fmt.Printf("admis")
	}else {
		fmt.Printf("non admis")
	}
}