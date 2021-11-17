<!--
ignore these words in spell check for this file
// cSpell:ignore Intn Errorf
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
> - Values can be of diffrent types
> - Keys don't support indexing
> - Used to represent a "thing" with a lot of diffrent properties
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

<!-- <details> -->
<summary>

</summary>

### Purpose of Interfaces
### Problems Without Interfaces
### Interfaces in Practice
### Rules of Interfaces
### Extra Interface Notes
### The HTTP Package
### Reading the Docs
### More Interface Syntax
### Interface Review
### The Reader Interface
### More on the Reader Interface
### Working with the Read Function
### The Writer Interface
### The io.Copy Function
### The Implementation of io.Copy
### A Custom Writer
### Assignment 2: Interfaces
### Assignment 3: Hard Mode Interfaces

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

<details>
<summary>
Stuff worth remembering
</summary>



### Go Types
- bool:  true or false. zero-value **false**
- string: text. zero=value **"" (empty string)**
- int: `int8`, `uint8` (byte), `int16`, `uint16`, `int32` (rune), `uint32`, `int64`, `uint64`, `int`, `uint`, `uintptr`. zero-value **0**
- float: `float32`, `float64`. zero-value **0**
- complex: `complex64`,`complex128` zero-value **0+0i**
- array
- slice: `[]<type>`
- map: `[<key type>]<value type>`

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
io |||
ioutil | import "io/util" | utility functions for IO | WriteFile,ReadFile
strings |import "strings"|basic string operations | Join, Split
os | import "os" |platform independent system calls | exit, Remove
time | import "time" |functionality for measuring and displaying time | Now, 
testing | import "testing" | tests package! | t.Error,  t.Errorf

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