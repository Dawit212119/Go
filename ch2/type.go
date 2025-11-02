package main

import "fmt"

func main(){
	var a int16=1000              // integer literal has base default of base 10 (Decimal)
	var b int8=1       
	var c int8=2          // integer literal has type default of int(depend on system arichtecture)
	fmt.Println(int8(a))         //   explict cast to smaller bit truncate the bits 
	fmt.Println(int16(b))         //      cast to higher size just add 0
	fmt.Println(b & c)
	fmt.Println(b << 2)

}