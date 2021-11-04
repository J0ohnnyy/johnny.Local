package main
import "fmt"
func main(){
	var n int
	fmt.Printf("tapez la valeur du dernier nombre:")
	fmt.Scanln(&n)
for i:=1; i<=n;i++{
		fmt.Println(i)
	}
}