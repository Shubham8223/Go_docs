package main

import "fmt"

// Variadic function to sum integers
func Sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    // Indexing example
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Second element:", nums[1]) // Output: 2

    // Spread example
    moreNums := []int{6, 7, 8}
    combined := append(nums, moreNums...) // Spread the second slice
    fmt.Println("Combined Slice:", combined) // Output: [1 2 3 4 5 6 7 8]

    // Rest parameters example
    total := Sum(1, 2, 3, 4, 5)
    fmt.Println("Total:", total) // Output: Total: 15
}
