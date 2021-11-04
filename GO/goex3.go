package main
import "fmt"
func main(){
	var a,b,c,m int
	fmt.Printf("tapez votre prmiere note:")
	fmt.Scanln(&a)
	fmt.Printf("tapez votre deuxieme note:")
	fmt.Scanln(&b)
	fmt.Printf("tapez votre troisiem note:")
	fmt.Scanln(&c)
	m =(a+b+c)/3
	fmt.Println("votre moyenne est:",m)
	fmt.Printf("montien:")
	switch {
	  case m<10 :
		fmt.Println("echoué.")
	  case m<12:
		fmt.Println("passabel.")
	  case m<14:
		fmt.Println("assez bien.")
	  case m<16:
		fmt.Println("bien.")
	  case m>16:
		fmt.Println("trés bien.") 		
	}
}