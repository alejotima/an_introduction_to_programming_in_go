/*
- All GO programs start with a package declaration
- Package are Go's way of organinizing and reusing code
- There are 2 types of programs in GO : Executables and libraries
- Executable are the kinds of programs that we can run directly from the terminal
- Libraries are collectionsof code
*/
package main // Package declaration

/*
- Import is a keyword for include code from other packages to use with our programs
- The fmt is the shorthand of "format"
*/
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
