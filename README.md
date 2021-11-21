<!--
ignore these words in spell check for this file
// cSpell:ignore Intn Errorf Holla PIMPL
-->

# udemy-Go-The-Complete-Developers-Guide

udemy course Go: The Complete Developer's Guide (Golang).

[Master the fundamentals and advanced features of the Go Programming Language (Golang)](https://www.udemy.com/course/go-the-complete-developers-guide/)


## Contents

1. [Getting Started](#getting-started)
2. [A Simple Start](#a-simple-start)
3. [Deeper Into Go](#deeper-into-go)
4. [Organizing Data With Structs](#organizing-data-with-structs)
5. [Maps](#maps)
6. [Interfaces](#interfaces)
7. [Channels and Go Routines](#channels-and-go-routines)
8. [Personal Takeaways](#other-takeaways)

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

<details>
<summary>
Creating a small project to get a better view of how Go does stuff.
</summary>

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

go uses the same seed by default. this is done with type Rand. we use `func New(src Source) *Rand`, which uses a Source type and returns a pointer type. then we can use it.


**deck.go**
```go
func (d deck) shuffle(){

    source := rand.NewSource(5)
    r := rand.New(source)

    for i := range d {
        newPosition :=r.Intn(len(d)-1)
        d[i],d[newPosition] = d[newPosition],d[i]
    }
}
```
to get a seed we will use the current time from the time package. tha main thing we learn is how to use the documentation.

```go
func (d deck) shuffle(){

    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    for i := range d {
        newPosition :=r.Intn(len(d)-1)
        d[i],d[newPosition] = d[newPosition],d[i]
    }
}
```

### Testing With Go

let's think about testing.

> "Go testing is not RSpec, mocha, jasmine, selenium, etc!"

we create a file ending with "_test.go", vscode will recognize this file and give us some quick actions. we run tests with `go test` from the command line.

### Writing Useful Tests

deciding what to test: we need to figure out what we care about, what we can test, etc...

we write our tests in functions, each function should test some functionality. we write test function with a capital letter, "TestNewDeck", then vscode shows us the option to run them directly, they all take a argument of type _\*testing.T_ called *t*.

`func TestSomethingName (t *testing.T)`, we use this parameter to create error messages and fail the test with `t.Error`, `t.Errorf`  and other functions.


> our first test:\
> when we create a new deck, before shuffling 
>       - it should be of a specific size
>       - the first card is "Ace of Spades"
>       - the last card is "Four of Clubs"


so our function starts like this

**deck_test.go**
```go
package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)
	if l != 16 {
		t.Errorf("Expected length 16, got %v", l)
	}
}

```

### Asserting Elements in a Slice

expanding the test, in a new deck, we know what the order the cards should be, so we can test this. we use the regular indexing notations.

**deck_test.go**
```go
func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)
	exp := 16
	if l != exp {
		t.Errorf("Expected length %d, got %v", exp, l)
	}
	// first := d[0]
	// last := d[l-1]
	expFirst, expLast := "Ace of Spades", "Four of Diamonds"
	first,last := d[0],d[l-1]
	if first != expFirst {
		t.Errorf("expected first card to be %v, was %v",expFirst, first)
	}
	if last != expLast {
		t.Errorf("expected last card to be %v, was %v",expLast, last)
	}
}
```


### Testing File IO

the final test is to save and read deck from disk. if we create a file and crush, we won't have an opportunity to clean it up! we have to take care of this.

> 1. delete any file in the current directory with the name "_deckTesting"
> 2. create a deck
> 3. save to file "_deckTesting"
> 4. load from file
> 5. assert deck length
> 6. delete any file in the current directory with the name "_deckTesting"


so how do we delete a file? the os package can help! `func Remove(name string) error` we delete before and after running the test



**deck_test.go**
```go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    filename := "_deckTesting"
    os.Remove(filename)
	d := newDeck()
	
    d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	l1 := len(d)
	l2 := len(loadedDeck)
	if l1 != l2 {
		t.Errorf("Expected length %d, got %v", l1,l2)
	}
    os.Remove(filename)
}
```

### Project Review

we used many packages, and we imported them all. we had a new type *deck*, built on top of slice of strings (*[]string*). we used receiver function and slices, which are more advanced then regular arrays.

the *deal* function wasn't a receiver function, it was to avoid ambiguity, the name might imply that we modify the original value. it's up to discussion. later on.

we also used the testing function, which has a type with a star, this will also be explained later.


### Assignment 1: Even and Odd


> In this assignment you'll get some practice with slices and for loops.
>
> Here are the directions:
>
> - Create a new project folder to house this new project and create a 'main.go' file inside of it.
> - In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
> - In the main function, create a slice of integers from 0 through 10.
> - Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
> - If the value is even, print out "even".  If it is odd, print out "odd".
> - Run your code from the terminal by changing into your project directory then running 'go run main.go'


</details>

## Organizing Data With Structs

<details>
<summary>
Creating Structs, initializing them. basic pointer operations.
</summary>

### Structs in Go

if we look back at the cards project, we used strings to denote the cards "Ace of Spades", how can we say that a card is from a specific suit or a specific value? we would have to do string operations. structs could help us with this, by using a struct to define a single card.

> struct :"data structure. collection of properties that are related together"

we will use a project "structs"

### Defining Structs

in the main.go file, we put the usual boiler plate code. our struct will define a person, with a first and last name (strings). we use the syntax of `type <name> struct`. we don't separate the fields, each in a new line

```go
package main

type person struct {
	firstName string
	lastName string
}

func main(){
	
}

```

### Declaring Structs

we have a type definition, now we need to instantiate this type. we can use positional arguments or name arguments with key:value pairs.

if we simply print a struct, we see the values inside curly braces, but not the field names

```go
	alex:= person{"Alex","Anderson"} //positional argument
	jack:= person{firstName:"jack",lastName: "jackson"} //names arguments

	fmt.Println(alex)
```

### Updating Struct Values

a final way to declare a struct is to declare a variable of that type, which will result in a struct with zero values in each field. so the person has empty strings as the values for first and last names. if we want to see things better, we can use the `fmt.Printf` syntax for a different view, if we pass the `%+v` as a format string, we can see the key:values pairs.

as most languages, we can update the values with the dot notation.

```go
	alex:= person{"Alex","Anderson"} //positional argument
	jack:= person{firstName:"jack",lastName: "jackson"} //names arguments
	var jon person // zero value initialization

	fmt.Println(alex,jack, jon)
	fmt.Printf("%+v\n",alex)
	fmt.Printf("%s %+s\n",alex.firstName,alex.firstName)
	jon.firstName = "jon"
	fmt.Println(jon)
```

### Embedding Structs

we can embed nest types inside other types. lets add a nested type to the person type. we also define a person with this information.

**in a multiline struct definition, we need a comma a the end of every line!**
```go

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main(){
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contact: contactInfo{
			email:   "party@jim.now",
			zipCode: 94000,
		},
	}

	fmt.Println(jim)
}
```

### Structs with Receiver Functions

when we create an embedded (nested) struct, we can omit the inner type member name, as use the typename as a field name as well. like javaScript es6 enhanced object literals. this will be handy later on.

```go
type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo //no member name
}
```

like with other types, we can create receiver functions for types, like the print function, and one function that changes the first name.

```go
func (p person) print(){
	fmt.Printf("%+v\n",p)
}
func (p person) updateName(newFirstName string){
	p.firstName =newFirstName
}
```
however, when we run this function, the name doesn't change! what happened?  time to learn about pointers...

### Pass By Value

quick refresher of RAM and memory.

go is pass-by-value (copy), not by reference, just like C. so if we want to modify a value in a function, we would need to use pointers.

### Structs with Pointers

first we update our code. then we'll talk about how

```go
func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName =newFirstName
}

jimPointer := &jim
jimPointer.updateName("jimmy")
jim.print()
```

### Pointer Operations

we have some new symbols,`&` and `*`. `&` is an operator to get a address of the variable. `*` is a pointer deference operator. get the value at that address. unless it's the type of a pointer.

so in the new updateName function, we changed the receiver type to be a pointer to a person.

> "turn address into value with `*address`"\
> "turn value into address with `&value`"

we will see more edge cases with pointers.

### Pointer Shortcut

instead of getting a Pointer and using it, we can use receiver function with a pointer type directly on a type. like taking the address and calling the correct function. this only works one way. we can't use value functions on pointers with this syntax.

```go
func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName =newFirstName
}

//jimPointer := &jim
//jimPointer.updateName("jimmy")
jim.updateName("jimmy") //also work
jim.print()
```

### Gotchas With Pointers


here is a possible gotcha. we create a slice and send it to a function by value, this should be a copy, so the value shouldn't change. but it does!

```go
package main

import "fmt"

func main() {
	mySlice := []string{"hi","There","How","Are","You?"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string){
	s[0]= "bye"
}
```

this is because of reference types and value types

### Reference vs Value Types

why did the value inside the slice change?

first of all, differences between slices and arrays: arrays are primitive data structure,they can't be resized, and we rarely use them. slices are used more often, they can grow and shrink. internally, they are built above arrays. 

a slice is a struct that manages the array, but it holds a pointer to the data, not the data itself. some types are value types (primitives, string, structs), and other are reference types (slices, maps, channels, pointers, functions). when we use reference types, we don't need to worry about taking pointers and stuff.

</details>

## Maps

<details>
<summary>
Go maps!
</summary>

### What's a Map?

maps are like dictionaries, they are statically types, so all keys must be of the same type, and the values must also all be from the same type(not the same type as the keys).

so we will start working with a new project, "maps". we can declare a map in several ways. one way is like this `<name> := map[<key type>]<value type>` and then start initializing the keys and the pairs. we need a trailing comma at the last pair

we will use mapping from colors to hex codes
```go
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff", //trailing comma!
	}
	fmt.Println(colors)
```

### Manipulating Maps

there are other ways to initialize maps, with the empty variable declaration or a `make(map[<key type>]<value type>)` function.\
we use the square brackets to manipulate the map, but we can't access records with the dot notation. we delete records with the `delete` function, which takes the map and the key.

```go
var colours map[string]int //zero values
otherColors := make(map[string]string) //also creates ma
otherColors["orange"]="badger"
delete (otherColors,"orange")
```

### Iterating Over Maps

lets make a function that takes a map, iterates over it and prints it. it's very similar to iterating over a slice, but we get the key and value rather than the index and value.

because map is a reference type, we don't need to send a pointer

```go
func printMap(c map[string]string){
	for color,hex := range c{
		fmt.Println("color",color,"hex code",hex)
	}
}
```

### Differences Between Maps and Structs

> Struct
> - Values can be of different types
> - Keys don't support indexing
> - Used to represent a "thing" with a lot of different properties
> - You need to know all the different fields at compile time
> - Value type
> Map:
> - All keys must be the same type
> - All values must be the same type
> - Keys are indexed, we can iterate over them
> - Used to represent a collection of related properties
> - Don't need to know all the keys at compile time
> - Reference type

if we change a map inside a function, this will change the values (reference types)

</details>

## Interfaces

<details>
<summary>
Using Interfaces for code reuse and other purposes.
</summary>

### Purpose of Interfaces

what interfaces do for us as developers.

>"We know that every value has a type, and every function has to specify the type of its arguments.\
so does that mean...\
Every function we ever write has to be rewritten to accommodate different types even if the logic is identical?"

for example, the deck shuffle function. do we care that it's called on a deck? does it use any special functions that use the deck, cards? NO! it uses the parts of the []string slice, but even then, it doesn't care about the string, only about the slice. this is one of the problems interfaces are meant to solve.

we'll start with an example, we'll write it naively and then fix it. we have a two chat-bots, one for english and one for spanish. they have many similarities. it wouldn't make sense to write the same code again. some function will have different code, but some are using those functions and will be identical. *getGreeting* will be different, but *printGreeting* will use the return value, but be the same. it would make sense to have one version of that function.

```go
type englishBot struct
func (englishBot) getGreeting() string
func printGreeting(eb englishBot)

type spanishBot struct
func (spanishBot) getGreeting() string
func printGreeting(sb spanishBot)
```

### Problems Without Interfaces

lets create the program, we'll use a new folder "interfaces". a note to remember is that if we aren't using a parameter, we can omit the name

```go
type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string{
	return "Holla!"
}
```
now we create the printGreeting functions, we get a re-declaration error! no function overloading!

```go
func printGreeting(eb englishBot){
	fmt.Println(eb.getGreeting())
}
func printGreeting(sb spanishBot){
		fmt.Println(sb.getGreeting())
}
```

how can we fix this problem? is there a way to pass the two different types to the same function?

### Interfaces in Practice

let's use interfaces. we add a new type, which is interface (not struct) and use it as the type of the function argument. `type <interface name> interface` and then we define requirements in terms of of function signatures.

```go
type bot interface{
	getGreeting() string
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}
```

this works!

interfaces are declared as requirement, a type doesn't need to declare itself to be part of an interface, it either is or it isn't based on satisfying the interface.

### Rules of Interfaces

here is another interface, with argument and return type. we can have more than one function

```go
type user struct {
	name string
}

type otherBot interface{
	getGreeting(string,int) (string, error)
	getBotVersion() float
	respondToUser(user) string
}
```

we have concrete types, which we can create values of, like primitives, slices, structs. we also have interface types, which are abstract ideas of types, we can't a instance of them.

### Extra Interface Notes

> - "Interfaces are **not** generic types". go doesn't support generic programming.
> - "Interfaces are implicit". we don't manually link types and interfaces (*duck typing?*)
> - "Interfaces are a contract to help us manage types".
> - "Interfaces are tough. Step #1 is understanding how to read them".

### The HTTP Package

using interfaces the standard library, we will start a folder "http". we'll look at the documentation to understand the types used in function `func Get (url string)(resp *Response, err error)`.

```go
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
```
for any other language, this would work, but not in go! where is the html body?

### Reading the Docs

we sent a request, but we didn't get an html body. we need to learn more. the response is a struct, it has a body field, with the type `io.ReadCloser` which is an interface type, but it doesn't show functions, it shows other interfaces!

```go
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read (p []byte)(n int, err error)
}

type Closer interface {
	Close() error
}
```

### More Interface Syntax

if we specify a interface as struct field, it's saying we don't care about the concrete type, just that it fulfills the interface. like the PIMPL idiom.

in go. we can define a interface as a composition of other interfaces.

```go
type interface A {
	foo()
}
type interface B {
	bar()
}
type interface C {
	A
	B
}
```

### The Reader Interface

let's understand the *Reader* interface. we have an imaginary world without interfaces, each function can return a different type, and for each type we need a different function with a different name and definition to handle them.

the Reader interface works as an adapter, we don't care about the type of the body, as long as it provides us the Reader interface. which populates a byte slice.

### More on the Reader Interface

understanding the reader interface. `Read(p []byte)(int, error)`. the function copies the raw data (unknown type and format) into the slice, and returns the number of bytes read and any errors encountered. because slice is a reference type, then it's possible to modify it inside the function

### Working with the Read Function

now lets try to read the data from the body. we declare a byte slice, pass it to the read function of the response Body element, and then cast it to a string to print.

```go
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	bs := make([]byte,99999) // create a slice with this size
	resp.Body.read(bs)
	fmt.Println(string(bs)) //byte slice is just like a string
}
```
however, go has built-in utilities to make this easier.

### The Writer Interface

here's an easier way to write the html to the console.

```go
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.stdout, resp.Body)
}
```

we used the Reader interface before, but we can also use Writer interface, which is the opposite. but this doesn't modify the data.

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

### The io.Copy Function

the Writer interface takes the data and writes it somewhere, the io.Copy function takes a reader and a writer, and uses them both together to copy data from one source to another.

io.Stdout is the of type _*File_, which has a function Write with the correct signature, so it fits the Writer interface.


### The Implementation of io.Copy

with the vscode editor, we can look at function implementations, we simply hold the <kbd>Ctrl</kbd> button and click on a function name, and we can see how it's implemented.

it creates a byte slice (if need), use a loop read and write data. if one of the values has *WriteTo* or *ReadFrom* functions, they are used instead.

```go
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOF or error.
// The return value n is the number of bytes read.
// Any error except EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to w until there's no more data to write or
// when an error occurs. The return value n is the number of bytes
// written. Any error encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```


### A Custom Writer

to satisfy the writer interface, our type needs `Write(p []byte) (n int, err error)`

so let's write such a type, as long as we have th correct function format, we can use the struct, even if the code is garbage.

```go
type logWriter struct {}

func (logWriter) Write(bs []byte) (n int, err error){
	return 1, nil
}
```
but lets be better
```go
func (logWriter) Write(bs []byte) (n int, err error){
	fmt.Println(string(bs))
	return len(bs),nil
}
```
### Assignment 2: Interfaces

> Write a program that create two custom struct types called 'triangle' and 'square.\
> The 'square' type should be a struct with a field called 'sideLength' of type float64.\
> The 'triangle' type should a struct with a field called 'height' of type float64, and a field called 'base' of type float64.
> 
> Both types should have function called 'getArea' that returns the calculated area of the square or triangle.\
> Area of triangle: 0.5 * base * height\
> Area of square: sideLength *sideLength
> 
> Add a shape interface that defines a function called 'printArea', this function should calculate the area of a the given shape and print it out to the terminal.\
> Design the interface so that 'printArea' function can be called with either a triangle or a square.

**note: the instructions are messed up, but it's pretty simple to understand what it wants us to do**

### Assignment 3: Hard Mode Interfaces


> Create a program the reads contents of a text file then prints its content's to the terminal. 
> 
> The file to open should be provided as an argument to the program when it's executed at the terminal. for example, `go run main.go myFile.txt` should open up the myFile.txt file.
> 
> To read in arguments provided to a program, you can reference the variable *os.Args* which is a slice of type sting. To open a file , check out yhe documentation for the *Open* functions in the *os* package.
> 
> What interfaces does the *File* type implement?\
> If the *File* type implements the *reader* interface you might be able to reuse the *io.Copy* function!

**maybe I should have closed the file...**

</details>

## Channels and Go Routines

<details>
<summary>
Communication between Go Routines using channels.
</summary>


### Website Status Checker

a program that checks the status of some popular websites, we will put this code in "channels" folder. we already know most of this.

```go
package main

func main() {
	links:= []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _,link := range links{

	}
}

```

### Printing Site Status

rather than put everything in the main function, let's create a different function to do this for
```go
func checkLink(link string){
	_,err:=http.Get(link)
	if err != nil{
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is responding")
}
```

this works, but not efficient, we are waiting for each request to finish before starting the next one.

### Serial Link Checking

could we skip this waiting time? why not start all request at the same time and print the responses in the order they return?

### Go Routines

goRoutines are one of the main features of go. the main program itself is a goRoutine. the http.Get function is a blocking call. so we want to fix this problem, this will be done by calling the function as a goRoutine.

To do this, we simply add the keyword `go` before the function call. this will create a goRoutine. once the goRoutine reaches a blocking call, it will yield control so the loop can continue.

there is more to know, and some edge cases.

### Theory of Go Routines

go has a Go Scheduler behind the scenes, when we have one cpu, the go scheduler monitors the go routines, and knows to yield cpu time according to io calls. go prefers to use one core, even if we have more.\
**(note: this isn't true today with more modern Go)**\
when we have multiple cpu cores, the scheduler tries to attach each goRoutine to a different cpu.

> "Concurrency is not parallelism"\
> Concurrency - we can have multiple threads executing code. if one thread blocks, another one is picker up and worked on.\
> Parallelism - Multiple threads executed at the exact same time, requires  multiple CPUs.

every program has a main goRoutine, the other routines are child routines, if the main routine ends, the rest of them also 'die'.

### Channels

now we need to use go routines, if we simply add the `go` keyword, then the function will end and we won't see any results!

```go
package main

func main() {
	links:= []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _,link := range links{
		go checkLink(link)
	}
	fmt.Println("finished with the loop!")
}
```

the problem is that the main routine ends, and this cancels the other routines. channels are used for communication between routines. channels are just like any other variable, they define the type of messages sent, like a string, float, a custom struct or booleans.

### Channel Implementation

channels are types, each channel has a type.

because channels are just like any other variable, if we want to use it in a function, we need to pass it to the function as an argument.

here we have some new syntax, 
```go
channel <- 5 // send value into channel
myNumber <-  channel // wait for value to be sent into the channel, and when it's done, assign the value 
fmt.Println(<- channel) // wait for the value to be sent into the channel, and use it as a function argument
```

```go
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down!"
		return
	}
	fmt.Println(link, "is responding")
	c <- "is responding!"
}
```

this time we get only one communication. it more than zero, but not what we wanted!

### Blocking Channels

the code that waits for a response from the channel is blocking, but when it receives data, it can continue.

if we put two statements that take data from the channel, we will see two sites.
if we put more "read" statements than the number of "write" statements we have, then we will hang.

### Receiving Messages

receiving a message is blocking call. instead of writing the statements one after another, we can write a regular for loop

```go
for i := 0; i < len(links); i++ {
	fmt.Println(<-c) //blocking code!
}
```

this works!

### Repeating Routines

one last change, we want to repeatedly ping each site, once it completes, we want to check this again! we will use the channel to return the link that we checked! we want to continuously, forever, check the links, so we modify the for loop.

```go
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println("out of loop!")
	for {
		fmt.Println(<-c) //blocking code!
	}
	go checkLink(<-c,c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is responding")
	c <- link
}
```
this works, but we are flooding the servers with requests, we should  wait between them.

### Alternative Loop Syntax

we wrote earlier this loop
```go
for {
	go checkLink(<-c,c)
}
```

but lets change the style to make it more clear. we move the link to the loop declaration, we use a range with a channel. 

```go
for l:= range c{
	go checkLink(l,c)
}
```

### Sleeping a Routine

let's add a pause between each call. we look at the documentation as usual, and we find *Sleep* inside the time package `Sleep (d Duration)`. *time.Second* is a duration type that we can do some arithmetics on.

But we need to decide where to pause the program. the sleep function operates on the **current goRoutine**.

if we put it in the main routine, things won't be great, because it needs to respond to message from other sites
```go
	for l := range c {
		time.Sleep( 5 * time.Second) // not here!
		go checkLink(l, c)
	}
```

### Function Literals

another option is to put the sleep in the checkLink function, but this will slow down the function 

```go
func checkLink(link string, c chan string) {
	time.Sleep( 5 * time.Second) // not here!
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is responding")
	c <- link
}
```

instead, we will put something called 'function literal', which is basically a lambda function.  we will also invoke it immediately. so the pause isn't part of the checkLink function, it's part of the block that defines the go routine.


```go
	for l := range c {
		go func(){
			time.Sleep(5 * time.Second)
			checkLink(l,c)
		}() //invoke
	}
```

but now we get some squiggly lines in the IDE

### Channels Gotcha!
   
we ge a warning that 
> "loop variable l is captured by function literal"

if we run the code, we get only one link called! this isn't what we wanted! so what's going on?

l is part of the outer scope, and is used in the inner scope. this is bad, so we need to do a pass by value instead, and we provide it to the function literal as an argument.


```go
	for l := range c {
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, c)
		}(l) //invoke
	}
```



</details>

## Other TakeAways

<details>
<summary>
Stuff worth remembering
</summary>

- go doesn't have function overloading - one name, one function
- the keyword `type`
  - is used to define type built on another type (alias, but with functions)
  - define structs
  - define interfaces
- no generics
- there is a var args (...) syntax
- go routines
  -  `go <func name>`
  -  `go func(){ /*body*/ }()` - function literal
  -  no `while` or `do{} while`, just `for` loops.
-  lambdas are called *function literals*
   -  function literals capture **by reference**, the values can be changed from either block.
-  writing to a channel is a blocking action.


### Go Types

- bool:  true or false. zero-value **false**
- string: text. zero=value **"" (empty string)**
- int: `int8`, `uint8` (byte), `int16`, `uint16`, `int32` (rune), `uint32`, `int64`, `uint64`, `int`, `uint`, `uintptr`. zero-value **0**
- float: `float32`, `float64`. zero-value **0**
- complex: `complex64`,`complex128` zero-value **0+0i**
- array
- slice: `[]<type>`
- map: `[<key type>]<value type>`
- structs: `type <struct name> struct`
- interfaces: `type <interface name> interface`
- channels: `make(chan <message type>`

**value types**: int, float, string, bool, structs.\
**reference types**: slices, maps,channels,pointers, functions.  


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
rand | import "math/rand" | randomness | IntN, newSource, New
debug |||
encoding |||
crypto |||
io |import "io"| IO primitives| Copy
ioutil | import "io/util" | utility functions for IO | WriteFile,ReadFile
strings |import "strings"|basic string operations | Join, Split
os | import "os" |platform independent system calls | exit, Remove, os.Args
time | import "time" |functionality for measuring and displaying time | Now, Sleep,
testing | import "testing" | tests package! | t.Error,  t.Errorf
http | import "net/http" | networking stuff | Get,Post

### Format String

how the string format works:
[formatting](https://zetcode.com/golang/string-format/)

> - d - decimal integer
> - o - octal integer
> - O - octal integer with 0o prefix
> - b - binary integer
> - x - hexadecimal integer lowercase
> - X - hexadecimal integer uppercase
> - f - decimal floating point, lowercase
> - F - decimal floating point, uppercase
> - e - scientific notation (mantissa/exponent), lowercase
> - E - scientific notation (mantissa/exponent), uppercase
> - g - the shortest representation of %e or %f
> - G - the shortest representation of %E or %F
> - c - a character represented by the corresponding Unicode code point
> - q - a quoted character
> - U - Unicode escape sequence
> - t - the word true or false
> - s - a string
> - v - default format
> - #v - Go-syntax representation of the value
> - T - a Go-syntax representation of the type of the value
> - p - pointer address
> - % - a double %% prints a single %

Type | Symbol | Example | Notes
-----|---------|---|---
decimal integers | %d| 0,5 | N/A
value | %v | any value
struct with field names | %+v | structs | prints the key:value pairs

</details>