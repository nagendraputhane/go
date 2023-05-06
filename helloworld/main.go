// how do we run the code in our project?
// go run main.go -> complies and executes
// go build -> just compile -> main.exe

// what does 'package main' mean?
// package = .go files. Package name at the beginning of code.
// package main is for executable package.
// Any other name is a reusable or library
// package main should have func main()
package main

//what does 'import "fmt"' mean?
//access to code in "fmt" package - part of standard library - format, I/O, print, scan
//"math", "encoding", "io", "crypto", etc..
//https://pkg.go.dev/std
import "fmt"

// what is func?
// func - function
// main package should have main func, which runs automatically
func main() {
	fmt.Println("Hi there")
}
