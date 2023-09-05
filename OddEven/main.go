package main
import "fmt"
type number []int
func main(){
	num := number{}
	for i:=1; i<=10;i++{
		num = append(num, i)
	}
	for number := range num {
		if number%2 == 0{
			fmt.Println("Even")
		} else{
			fmt.Println("Odd")
		}
	}
}