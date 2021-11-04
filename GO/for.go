package main
import "fmt"
func main(){
	var x,z int
	fmt.Printf("donnez la limite:")
    fmt.Scanln(&x)
	fmt.Printf("De:")
	fmt.Scanln(&z)
	for i:= z; i<=x;i++{
		fmt.Println(i)
	}
}