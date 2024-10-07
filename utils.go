package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "reflect"
    "strings"
    "time"
)

func main() {
    // 1. Basic Functions
    arr := []int{1, 2, 3}
    fmt.Println("Length:", len(arr)) // Output: Length: 3

    slice := make([]int, 5, 10)
    fmt.Println("Length:", len(slice), "Capacity:", cap(slice)) // Output: Length: 5 Capacity: 10

    // 2. String Functions
    str := "Hello, World!"
    fmt.Println("Contains 'World':", strings.Contains(str, "World")) // Output: true
    fmt.Println("Has Prefix 'Hello':", strings.HasPrefix(str, "Hello")) // Output: true
    fmt.Println("Has Suffix '!':", strings.HasSuffix(str, "!")) // Output: true
    fmt.Println("Index of 'o':", strings.Index(str, "o")) // Output: 4
    fmt.Println("Split by ',':", strings.Split(str, ",")) // Output: [Hello  World!]
    fmt.Println("To Upper:", strings.ToUpper(str)) // Output: HELLO, WORLD!
    fmt.Println("To Lower:", strings.ToLower(str)) // Output: hello, world!
    fmt.Println("Trimmed:", strings.TrimSpace("   Hello   ")) // Output: Hello

    // 3. Math Functions
    x := -3.5
    fmt.Println("Absolute:", math.Abs(x)) // Output: 3.5
    fmt.Println("Power (2^3):", math.Pow(2, 3)) // Output: 8
    fmt.Println("Square Root of 16:", math.Sqrt(16)) // Output: 4
    fmt.Println("Sin(π/2):", math.Sin(math.Pi/2)) // Output: 1

    // 4. Collection Functions
    slice = append(slice, 4, 5)
    fmt.Println("Appended Slice:", slice) // Output: [0 0 0 0 0 4 5]

    src := []int{1, 2, 3}
    dst := make([]int, len(src))
    copy(dst, src)
    fmt.Println("Copied Slice:", dst) // Output: [1 2 3]

    // 5. Channel Functions
    ch := make(chan int, 2) // Buffered channel
    ch <- 1
    ch <- 2
    fmt.Println("Channel Length:", len(ch)) // Output: 2
    close(ch) // Close the channel

    // 6. Error Handling
    err := fmt.Errorf("this is an error")
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: this is an error
    }

    // 7. Type Conversion Functions
    var y float64 = 3.14
    var z int = int(y)
    fmt.Println("Converted int:", z) // Output: Converted int: 3

    // 8. Reflection Functions
    var a interface{} = "Hello, Go!"
    fmt.Println("Type of a:", reflect.TypeOf(a)) // Output: Type of a: string

    // 9. File and I/O Functions
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    file.WriteString("Hello, World!")

    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("File Content:", string(content)) // Output: Hello, World!

    // 10. JSON Functions
    data := map[string]string{"name": "Alice", "city": "Wonderland"}
    jsonData, _ := json.Marshal(data)
    fmt.Println("JSON:", string(jsonData)) // Output: JSON: {"city":"Wonderland","name":"Alice"}

    var result map[string]string
    json.Unmarshal(jsonData, &result)
    fmt.Println("Decoded:", result) // Output: Decoded: map[city:Wonderland name:Alice]

    // 11. Time Functions
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime)

    formattedTime := currentTime.Format("2006-01-02 15:04:05")
    fmt.Println("Formatted Time:", formattedTime)

    parsedTime, _ := time.Parse("2006-01-02", "2024-10-07")
    fmt.Println("Parsed Time:", parsedTime)
}
