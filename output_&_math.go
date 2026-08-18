// code structure in GO

package main
import "fmt"

func main(){
	// your code goes here
}


// print command in GO

package main
import "fmt"

func main() {
	fmt.Print(12)
}


// "I Love GO" program

package main
import "fmt"

func main() {
    fmt.Print("I Love GO") 
}


// arithmetic operators in GO

package main
import "fmt"

func main() {

    fmt.Print(13+3)
    fmt.Print(13-3)
    fmt.Print(13*3)
    fmt.Print(13/3)
}


// multiple outputs in GO

package main
import "fmt"

func main() {
	fmt.Print(3 + 4)
	fmt.Print(2 + 1)
}


// output in separate lines

package main
import "fmt"

func main() {

    fmt.Println(3 + 4)
    fmt.Println(2 + 1)
}


// area and perimeter of rectangle

package main
import "fmt"

func main() {
    fmt.Println(11 * 13)         
    fmt.Println(2 * (11 + 13))
}


// inserting space between outputs

package main
import "fmt"

func main() {
    fmt.Print(3 + 4)
    fmt.Print(" ")
    fmt.Print(2 + 1)
}


// text between outputs

package main
import "fmt"

func main() {
    fmt.Print(3 + 4)
    fmt.Print(" and ")
    fmt.Print(2 + 1)
}


// extra indenting cause no error - example 1

package main
import "fmt"

func main() {
	fmt.Print(12)
        fmt.Print(11)
}


// extra indenting cause no error - example 2

import "fmt"

func main() {
	fmt.Print(12)

        fmt.Print(11)
}


// declrare variable - explicit variable declaration

import "fmt"
func main() {
	var age int := 30
	var name string := "John Doe"
	var temperature float64 := 25.5
	fmt.Print(age)
	fmt.Print(name)
	fmt.Print(temperature)
}


// type inference - explicit variable declaration

import "fmt"
func main() {
	age := 30
	name := "John Doe"
	temperature := 25.5
	fmt.Print(age)
	fmt.Print(name)
	fmt.Print(temperature)
}
