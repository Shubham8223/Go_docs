package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
	City *string `json:"city"`
}

func main() {
	city := "Pune"
	p := Person{Name: "Shubham", Age: 25, City: &city}

	// 1. Print type and value
	fmt.Println("Type:", reflect.TypeOf(p))
	fmt.Println("Value:", reflect.ValueOf(p))

	// 2. Loop over fields and print names + values
	val := reflect.ValueOf(p)
	typ := reflect.TypeOf(p)
	fmt.Println("\nStruct fields:")
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("Field: %s | Value: %v | JSON Tag: %s\n", field.Name, value.Interface(), field.Tag.Get("json"))
	}

	// 3. Check if a pointer is nil
	var ptr *int
	fmt.Println("\nCheck if pointer is nil:", isNil(ptr))

	// 4. Update a struct field using reflection
	book := &Book{Title: "Go Basics", Price: 299}
	updateField(book, "Title", "Mastering Go")
	updateField(book, "Price", 499)
	fmt.Println("\nUpdated Book:", book)

	// 5. Type checking
	checkType([]string{"go", "rust"})
	checkType(Person{})
	checkType(42)
}

// Check if an interface is nil or nil pointer
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Ptr && v.IsNil()
}

// Struct to use for updating fields
type Book struct {
	Title string
	Price int
}

// Dynamically update struct fields by name
func updateField(obj interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByName(fieldName)
	if f.IsValid() && f.CanSet() {
		val := reflect.ValueOf(value)
		if f.Type() == val.Type() {
			f.Set(val)
		}
	}
}

// Detect if input is a struct/slice/other
func checkType(i interface{}) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Struct:
		fmt.Println("✅ Struct type detected.")
	case reflect.Slice:
		fmt.Println("✅ Slice type detected.")
	default:
		fmt.Println("🌀 Other type:", v.Kind())
	}
}
