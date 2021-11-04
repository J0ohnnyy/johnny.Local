package  main 

import "fmt"
func main() {
	   var h,m,s, totalensecond int
	fmt.Printf("donnez nombre d'heure:")
	fmt.Scanln(&h)
	fmt.Printf("donnez nombre de minute:")
	fmt.Scanln(&m)
	fmt.Printf("donnez nombre de seconde:")
	fmt.Scanln(&s)
	totalensecond = h * 3600+ m * 60 + s
	fmt.Println("le total en second:", totalensecond)
}