package main

import "fmt"

// ----------------------------------------
// 1. Encapsulation
// ----------------------------------------

// Person represents an individual with public and private fields
type Person struct {
    Name string // Public field
    age  int    // Private field
}

// SetAge sets the age of the person (public method)
func (p *Person) SetAge(age int) {
    p.age = age
}

// GetAge retrieves the age of the person (public method)
func (p *Person) GetAge() int {
    return p.age
}

// ----------------------------------------
// 2. Inheritance (via Composition)
// ----------------------------------------

// Address represents a location
type Address struct {
    City    string
    Country string
}

// User represents a user with an embedded Address struct
type User struct {
    Name    string
    Address // Composition: User has an Address
}

// ----------------------------------------
// 3. Polymorphism
// ----------------------------------------

// Shape interface defines a method for calculating area
type Shape interface {
    Area() float64
}

// Circle struct implements the Shape interface
type Circle struct {
    Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Rectangle struct implements the Shape interface
type Rectangle struct {
    Width, Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// PrintArea function takes a Shape interface and prints its area
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

// ----------------------------------------
// 4. Abstraction
// ----------------------------------------

// Animal interface defines a behavior
type Animal interface {
    Speak() string
}

// Dog struct implements the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

// Cat struct implements the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

// MakeSound function takes an Animal interface and makes it speak
func MakeSound(a Animal) {
    fmt.Println(a.Speak())
}

// ----------------------------------------
// Main function to run examples
// ----------------------------------------
func main() {
    // Encapsulation example
    p := Person{Name: "Alice"}
    p.SetAge(30)
    fmt.Println(p.Name, p.GetAge()) // Output: Alice 30

    // Composition example
    user := User{
        Name: "Bob",
        Address: Address{
            City:    "New York",
            Country: "USA",
        },
    }
    fmt.Println(user.Name, user.City, user.Country) // Output: Bob New York USA

    // Polymorphism example
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
    PrintArea(c) // Output: Area: 78.5
    PrintArea(r) // Output: Area: 24

    // Abstraction example
    d := Dog{}
    cat := Cat{}
    MakeSound(d) // Output: Woof!
    MakeSound(cat) // Output: Meow!
}
