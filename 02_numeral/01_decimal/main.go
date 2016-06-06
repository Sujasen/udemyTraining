package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Printf("%d - %b - %x - %#x \n", 42, 42, 42, 42)
	fmt.Printf("%s \t %s \t %s \t %s \n","Dec" ,"Bi" ,"Hex", "UTF8")
	for i := 0; i <= 4; i++{
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
