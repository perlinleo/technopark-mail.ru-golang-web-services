package main

import (
	"fmt"
	"log"
	c "mycalc/calc"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		log.Fatal("No expression given");
	}

	expr := os.Args[1]

	result, err := c.Calc(expr)
	if err!=nil{
		fmt.Println("Error: ",err);
		return;
	}
	fmt.Println(result);
	
}