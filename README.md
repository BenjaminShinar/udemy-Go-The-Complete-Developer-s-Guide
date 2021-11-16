<!--
ignore these words in spell check for this file
// cSpell:ignore Intn
-->

# udemy-Go-The-Complete-Developers-Guide

udemy course Go: The Complete Developer's Guide (Golang).

[Master the fundamentals and advanced features of the Go Programming Language (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)


## Contents

1. Getting Started
2. A Simple Start
3. Deeper Into Go
4. Organizing Data With Structs
5. Maps
6. Interfaces
7. Channels and Go Routines
8. Extras

[course diagrams](https://github.com/StephenGrider/GoCasts/tree/master/diagrams)

## Getting Started
<details>
<summary>
getting go, vscode, vscode go extensions.
</summary>
[go](https://go.dev/)\
[vscode](https://code.visualstudio.com/)\
[vscode go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
</details>


## A Simple Start
<details>
<summary>
Basic Go program
</summary>

### Boring Ol' Hello World

create a "hello world" folder, with "main.go" file

```go
package main

import "fmt"

func main(){
    fmt.Println("hi there!")
}
```

### Five Important Questions

the five basic questions that we can learn from the simple example

> 1. How do we run the code in our project?
> 2. What does `package main` mean?
> 3. What does `import "fmt"` mean?
> 4. What's that `func` thing?
> 5. How is the file "main.go" organized?

now we want to run this code, we go to our folder, and we can run the code

```sh
cd hello world
go run main.go
```

but what's with that `go` command? what else can it do?

> - go build  - compiles a bunch of go source code files
> - go run - compiles and executes one or two files
> - go fmt - formats all the code in each file in the current directory
> - go install - compiles and "installs" a package
> - go get - downloads the raw source code of someone elses package
> - go test - runs any tests associated with the current project

we first run `go run`, but if we run `go build`, it will create an executable file.

### Go Packages

in go:
> Package == Project == Workspace

a package can have many files, each files declares which package it belongs to. in our example, `package main`.

there are two types of packages, executables and reenables. executables packages create executable files, reusable packages are support libraries (dependency).

executables packages are always called 'main', and should always have a `func main()`. any other name makes it a support package, and will not create an executable package.

if we change the name of the package and build it, it won't create an executable file.

### Import Statements

the import statement gives us access to function defined in another package. some packages are part of the standard, and some are not.

[packages documentation](https://pkg.go.dev/std)

### File Organization

`func` is short for function. the required format is `func`, the name of the function, arguments to the function, curly braces (same line), and then the body in a new line, and the closing curling braces in a new line.

every file in this course will look like this:
> - package declaration
> - imports other packages that we need
> - body: define functions, things that we want to do


</details>

## Deeper Into Go

### Project Overview
we start with a project: "cards", this project will be about playing card games, we will create the following functions:

> - newDeck - create a list of playing card. essentially an array of string
> - print - log out the contents of a deck of cards
> - shuffle - shuffle all the cards in a deck
> - deal - create a 'hand' of cards.
> - saveToFile - save a list of cards to a file on the local machine
> - newDeckFromFile - load a list of cards from the local machine

### New Project Folder

all of our files will be inside a single directory, so we need 'cards' folder. everything for the project will be inside this folder. we start with another main.go file.

```go
package main

func main() {

}
```

### Variable Declarations

now we get into some actual parts of go, we will start with variables declarations. there are some ways to define variables;

1. `var <var name> <type> = <value>` - has a warning message.
   1. var informs go that this is a new variable
   2. next is the name of the variable
   3. then the type
   4. the equal sign
   5. the initial value
   6. can live outside a function
2. `<var name> := <value>` - type inference
   1. `:=` creates a new variable. 
   2. must provide an initial value
   3. must be inside a function

go is a static typed language (like c++ and java, unlike javascript, python or ruby). if a variable is from one type, then it can't be changed to hold a value of another type.

types:
- bool
- string
- int
- float64
- (some more)

```go
package main

import "fmt"

//card4 := 55 //illegal - can't do this outside
var card3 int = 60 legal - can do this outside

func main() {
    //option 1
    var card1 string = "Ace of Spades"
    fmt.Println(card1)
    //option 2
    card2 := "Ace of Spades"
    fmt.Println(card2)

    fmt.Println(card3)
}
```

### Functions and Return Types

defining a function is the same as the "main" function, but while the main function didn't return anything, we want our function to return a value.

`func <name>(<arguments>) <return type>` then `{` on the same line, the function body in a new line and the closing `}` on a new line again. when we don't return anything, we can omit the return type.

type inference works with functions without issues.

```go
package main

import "fmt"

func main() {
    var card1 string = newCard()
    fmt.Println(card1)
}

func newCard() string{
    return "Five of Diamonds"
}
```

### Slices and For Loops

Go has two types to handle 'lists' of data: Array and Slice. an array is a fixed length list, a slice can grow and shrink. for both options the type of each record must be the same.

to create a slice, we use the following syntax: `<name> := []<type>{<value1>,<value2>}`. we can add elements with the `append(<slice>, <value>)` function. this function doesn't modify the original slice, it returns a new one.

we can iterate over slice with a `for <index>,<element> := range <slice> {` syntax. then we can define what we do inside. notice that we use the `:=` syntax. we can also define them outside and use `=` instead.

```go
package main

import "fmt"

func main() {
	cards1 := []string{"Ace of Diamonds", newCard()}
	cards2 := append(cards1, "Five of Hearts")
	fmt.Println(cards1)
	fmt.Println(cards2)
	for i, card := range cards2 {
		fmt.Println(i, card)
	}

    var i2 int
	var card2 string
	for i2, card2 = range cards1 {
		fmt.Println(i2, card2)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
```

### OO Approach vs Go Approach

Go isn't a a object oriented language. there are no classes (there are structs). if we were using an OOP approach, we would have a Deck class that we can initialize Deck instances from.

we will define a Deck type, which is built over `[]string` (slice of string), and define functions that use this type as a "receiver".

we will also create other files "deck.go" and "deck_test.go" to hold the type definition and some associated test.

### Custom Type Declarations

we define a type on top of an existing type, but this allows us to use it in functions. `type <name> <base type>`. we can now replace `[]string` with the newly created `deck`.

**deck.go**
```go
package main

//create a new type 'deck' which is a slice of strings

type deck []string
```
we also add the other file to the `go run` command.
```sh
go run main.go deck.go
```

now we also want a function that does the print loop for the deck of cards
```go
package main

//create a new type 'deck' which is a slice of strings

type deck []string

func (d deck) print (){
    for i,card:= range d {
        fmt.Println(i,card)
    }
}
```

now we can call this function in the main.go file, using the dot syntax.


### Receiver Functions

a receiver function is sort of like a method, or an extension method ( like *this*, or *self*)

`func (<identifier> <type>) <function name> (<arguments>) <return type>`

by convention we use a one or two letter name to refer to the passed value, and usually the abbreviation of the type.
**deck.go**
```go
package main

//create a new type 'deck' which is a slice of strings

type deck []string

func (d deck) print (){
    for i,card:= range d {
        fmt.Println(i,card)
    }
}
```


**main.go**
```go
package main

import "fmt"

func main() {
	cards := deck{"ace of Diamonds", "aaa"}
    cards.print()
}
```


### Creating a New Deck

out of the six functions we set out to create, we have just created one! we will now create another function, *newDeck()*. this is not a receiver function, it creates a new deck.

here we also see a nested loop. we usually need to define a index and the value in range for loop. so we use the `_` instead to make this a throw-away variable.

**deck.go**
```go
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}
```

### Slice Range Syntax

now we want to work on the "deal" function, which will create a "hand" (still a deck) of some size. go is zero-indexed, and uses the square brackets as an indexing operator. we can use a helper syntax to take a range, starting from an index, and going up to an ending index (not included). we can omit one of the indexes and that would mean the beginning or the end (including).


```go
	numbers := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(numbers[:3])  //0,1,2
	fmt.Println(numbers[3:])  //3,4,5,6
	fmt.Println(numbers[4:6]) //4,5
```


### Multiple Return Values

now lets implement the *deal* function, we define the function and the arguments, the identifier comes before the type name (unlike other languages), we want to return two separate values. we already saw something similar with the range for loops. so we define the return types in parentheses. when we return them we stick a comma.

**deck.go**
```go
func deal (d deck, handSize int) (deck,deck) {
    return d[:handSize], d[handSize:]
}
```

we get the variables returned from the function with a comma, just like what we used for the range loops. we can use either `:=` syntax to create new variables, `=` to assign, and `:=` when we assign one and create another
```go
	cards := newDeck()
	cards.print()
	h1, h2 := deal(cards, 2)
	fmt.Println("hand 1")
	h1.print()
	fmt.Println("hand 2")
	h2.print()
```

### Byte Slices

now we will look at the "saveToFile" function. we want to save the data to disk. we will use another package from the go standard, this time "ioutil" (io utilities), and the *WriteFile* function.

> `func WriteFile(filename string, data []byte,perm os.FileMode) error`

we can see that this function uses a non basic type, *os.FileMode*, and returns type *error*. to use the function we need to turn the deck into a []byte - a slice of bytes. a string is basically a byte slice, each byte is an ascii code.

[ascii table](https://www.asciitable.com)

### Deck to String

we want to make our deck into a string and then a []byte (slice of bytes). this will require **type conversion**, this is done with using the type like an operator.

```go
sliceOfBytes := []byte("hello")
fmt.Println(sliceOfBytes) //[104 101 108 108 111]
```

we have a deck, which is []string, we want it to be a single string, and the make that string into []byte. we might also have other cases where we need a single string from a deck, so lets create a function for it, and then use it. we can make it receiver function or a free function, we will go with a receiver function syntax.

**deck.go**
```go
func (d deck) toString() string {

}
```

### Joining a Slice of Strings

the deck type is []string, so we want to make a single string out of it. we start with type conversion. this is safe to do. we will use another package that can help us with this, 

**deck.go**
```go
import "strings"
func (d deck) toString() string {
	s := []string(d)
	str := strings.Join(s, ",")
	return str
    //return strings.Join(d, ",") // also works
}
```

we have multiple imports, so we can condense them together, we use parentheses and each package on a new line (no separator)
```go
import (
    "fmt"
    "strings"
)
```

now we have single string, so we are almost ready.

### Saving Data to the Hard Drive

now we need the final 'saveToFile' function, it will be a receiver function, which will make use of ioutil's *WriteFile*, and will return the same type of value. we will use 0666 as the permissions,


```go
import "io/util"

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
```
we can test this. if we pass a simple string this will be treated as a relative path, and will be created at same directory as the executable.

**main.go**
```go
d := newDeck()
d.saveToFile("abc.txt")
```

### Reading From the Hard Drive

now we do things in reverse, let's find the matching function in ioutil package. this is `func ReadFile(filename string)([]byte,error)`. for now we care about the byte slice part.

but what is the error object? if there was an error, it will have a value, otherwise, the value will be **nil**. we need to understand the error and decide how to deal with it. we can return a new deck, or decide that something we wrong and quit the program. we decide that in this case, we should quit the program.
we will use the os (operating system) package and exit with status code 1.

**deck.go**
```go
func newDeckFromFile(filename string) deck{
    bs,err= ioutil.ReadFile(filename)
    if err != nil {
        // Option #1 - log the error and return a new deck
        // Option #2 - log the error and quit the program
        fmt.Println("Error:",err)
		os.Exit(1) // exit with status code

    }

   // return something
}
```

### Error Handling

if there was no error, we have a []byte, which we want to make into a string, and then []string, eventually a deck. we will use type conversions and *strings.Split()*. the go language convention is to use short variables names, even single letter.

**deck.go**
```go
func newDeckFromFile(filename string) deck{
    bs,err= ioutil.ReadFile(filename)
    if err != nil {
        // Option #1 - log the error and return a new deck
        // Option #2 - log the error and quit the program
        fmt.Println("Error:",err)
		os.Exit(1) // exit with status code

    }

   return deck(strings.Split(string(bs),","))
}
```

we can test this in the main file.

### Shuffling a Deck

now we want to shuffle the deck, there is no built in shuffle functionality, but we can do one ourselves. we iterate over the deck, and for each card, we randomly swap it with a different card.

we need the length of the deck and a way to generate a random number, this comes from "math/rand", the function `func Intn (n int) int`. we find the length of the slice with *len*. we use the range for loop syntax, but drop the 2nd argument. and now we use the fancy way to swap.

**deck.go**
```go
func (d deck) shuffle(){
    for i := range d {
        newPosition :=rand.Intn(len(d)-1)
        d[i],d[newPosition] = d[newPosition],d[i]
    }
}
```
**main.go**
```go
	cards5 := newDeck()
	cards5.shuffle()
	cards5.print()
```

we run this several times, and we always get the same result! we didn't set up a random seed value!
### Random Number Generation
### Testing With Go
### Writing Useful Tests
### Asserting Elements in a Slice
### Testing File IO
### Project Review
### 1: Even and Odd



## Organizing Data With Structs

<details>
<summary>

</summary>
</details>

## Maps

<details>
<summary>

</summary>
</details>

## Interfaces

<details>
<summary>

</summary>
</details>

## Channels and Go Routines

<details>
<summary>

</summary>
</details>

## Extras

<details>
<summary>

</summary>
</details>
   
## Other TakeAways



### Go Types
- bool:  true or false.
- string: text
- int: `int8`, `uint8` (byte), `int16`, `uint16`, `int32` (rune), `uint32`, `int64`, `uint64`, `int`, `uint`, `uintptr`.
- float: `float32`, `float64`.
- complex: `complex64`,`complex128`
- array
- slice
- map
  
### Go CLI

- *go build*
- *go run*
- *go fmt*
- *go install*
- *go get*
- *go test*
  
### packages used

Package name | Import statement | Usage | Noticeable functions
-------|--------------|-----------|----
fmt, format | import "fmt"' | basic IO to console with formatting |Println,
math | ||
rand | import "math/rand" | randomness | IntN
debug |||
encoding |||
crypto |||
io |||
ioutil | import "io/util" | utility functions for IO | WriteFile,ReadFile
strings |import "strings"|basic string operations | Join, Split
os | import "os" |platform independent system calls | exit