package main //  the main package .  within a Go module, code is organized into one or more packages.
//   the main package in a Go module contains the code that starts a Go program. which is the main function

//  go is compiled language, compiled anywhere, run on the target

import "fmt" //   unlike js,java,c   in go import  the whole package , no individual functions,constant,types

func main() {
	var i int
	fmt.Println("Hello, world!")
    fmt.Println(i)
	fmt.Println(0b001);
	fmt.Println(0o001011)
	fmt.Println(0o00010010)
	fmt.Println(0b00010010)
  fmt.Println(6.03e23)
  fmt.Println(141);
  fmt.Println('\141')
  fmt.Println('a')
  fmt.Println('\n')

  fmt.Println("hello'a's")
fmt.Printf("hello %d\n",'a')
	fmt.Printf("Hello, %s\n", "dawit")
	fmt.Printf("%c\n", '\141')
	var x uint8=100;
	var y byte =0b00000001;
	fmt.Println(x+y);
}

///   go fmt to format the  code
// go vet  to scan for possible bugs
