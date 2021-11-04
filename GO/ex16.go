package main 
import "fmt"
func main(){
	var H,M int
	fmt.Printf("tapez la valeur d'heure:")
	fmt.Scanln(&H)
	fmt.Printf("tapez la valeur de minute:")
    fmt.Scanln(&M)
	 M=M+1
	if M==60 {
		M=0
		H=H+1
		if H==24{
			H=00 }
	}  
	fmt.Println("L'heure apres une minute:",H,":",M)
}	