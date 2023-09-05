package main
import "fmt"

func main(){

	colors := map[string]string{
		"red":"ff0000",
		"green":"ff3887",
	}

	// var colors map[string] string
	// colors := make(map[string]string)
	colors["white"] = "#ffffff"

	// delete(colors, "white")
	printMap(colors)
	// fmt.Println(colors)
}
func printMap(m map[string]string){

	for key, value := range m{
		fmt.Println(key, value)
	}
}