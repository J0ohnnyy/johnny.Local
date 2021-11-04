package main
import "fmt"
func main(){
	var n int
	fmt.Printf("donnez un nombre entre 1 et 3:")
	fmt.Scanln(&n)
	for n<1 || n>3{
		fmt.Printf("donnez un nombre entre 1 et 3:")
	fmt.Scanln(&n)
	}
}