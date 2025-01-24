// packages/basics/main.go
package main

//import  the fmt package from the standard library
import (
	"fmt"
	"time"
)

func main() {
	//user the fmt package to print the string "Hello Gopher"
	fmt.Println("Hello Gopher!")

	// use the time package to print the current workday
	fmt.Printf("Today is %s", time.Now().Weekday())
}
