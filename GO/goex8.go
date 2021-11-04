package main
import "fmt"
func main(){
	var x,y,m float64
	fmt.Printf("donnez la nombre d'heure du travaille:")
	fmt.Scanln(&x)
	fmt.Printf("donnez le salaire horaire:")
	fmt.Scanln(&y)
	switch {
		case x<=40:
			m=x*y
			fmt.Println("votre montant sera:",m)
		case x<=45:
			m=40*y+(x-40)*y*1.1
			fmt.Println("votre montant sera:",m)
        case x<=50:
             m=40*y+5*y*1.1+(x-45)*y*1.25
			 fmt.Println("votre montant sera:",m)
		case x>50:
			m=40*y+5*y*1.1+5*y*1.25+(x-50)*y*1.5
			fmt.Println("votre montant sera:",m)	   
	}
}