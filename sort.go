package main

import (
    "fmt"
    "sort"
)

// Person struct
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    // Create a slice of Person
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    fmt.Println("Before sorting:", people)

    // Sort by age
    sort.Sort(ByAge(people))

    fmt.Println("After sorting by age:", people)
}
