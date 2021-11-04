package main
import "fmt"
func main(){
	var  nbr,max,pos,i int
	fmt.Printf("Donnez un nombre:")
	fmt.Scanln(&nbr)
	max=nbr
	pos=1
	i=1
	for nbr!=0{
		if nbr>max {
			max=nbr
			pos=i
		}
		
	  
		      fmt.Printf("Donnez un nombre:")
		      fmt.Scanln(&nbr)
			  i++
		   
		    }
	    
    
    fmt.Println("Le nombre le plus grand est :",max)
	fmt.Println("La position du plus grand nombre est:",pos)
}