package main 
import "fmt"
func main(){
	var H,M,S int
	fmt.Printf("tapez la valeur d'heure:")
	fmt.Scanln(&H)
	fmt.Printf("tapez la valeur de minute:")
    fmt.Scanln(&M)
	fmt.Printf("tapez la valeur de seconde:")
	 S=S+1
	if S==60{
        S=0
		M=M+1
	}if M==60 {
		M=0
		H=H+1
		if H==24{
			H=00 }
	}  
	fmt.Println("L'heure apres une seconde:",H,":",M":",S)
}	