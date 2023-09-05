package main
import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}
	printGreeeting(eb)
	printGreeeting(sb)
}
func (englishBot) getGreeting ()string{
	//custom logic 
	return "Hi There!"
}
func (spanishBot) getGreeting ()string{
	return "Hola !!"
}
func  printGreeeting(b bot){
	fmt.Println(b.getGreeting())
}