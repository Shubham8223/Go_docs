package main

import "fmt"

// ----------------------------------------
// 1. Arrays
// ----------------------------------------

// Arrays are fixed-size collections of elements of the same type.
var arr [5]int

// Example of using an array
func arrayExample() {
    arr[0] = 1
    arr[1] = 2
    fmt.Println(arr) // Output: [1 2 0 0 0]
}

// ----------------------------------------
// 2. Slices
// ----------------------------------------

// Slices are dynamic-size, flexible views into arrays.
func sliceExample() {
    slice := []int{1, 2, 3} // Creating a slice
    slice = append(slice, 4) // Adding an element
    fmt.Println(slice) // Output: [1 2 3 4]
}

// ----------------------------------------
// 3. Maps
// ----------------------------------------

// Maps are unordered collections of key-value pairs.
func mapExample() {
    m := make(map[string]int) // Creating a map
    m["one"] = 1
    m["two"] = 2
    value := m["one"] // Accessing a value
    fmt.Println(value) // Output: 1
}

// ----------------------------------------
// 4. Structs
// ----------------------------------------

// Structs allow you to create custom data types.
type Person struct {
    Name string // Name of the person
    Age  int    // Age of the person
}

// Example of using a struct
func structExample() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p) // Output: {Alice 30}
}

// ----------------------------------------
// 5. Interfaces
// ----------------------------------------

// Interfaces define a set of methods that types must implement.
type Shape interface {
    Area() float64 // Method to calculate area
}

// Circle struct implements Shape interface
type Circle struct {
    Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Example of using an interface
func interfaceExample() {
    var s Shape = Circle{Radius: 5}
    fmt.Println(s.Area()) // Output: 78.5
}

// ----------------------------------------
// 6. Linked List
// ----------------------------------------

// A linked list consists of nodes, where each node points to the next.
type Node struct {
    Value int   // Value of the node
    Next  *Node // Pointer to the next node
}

type LinkedList struct {
    Head *Node // Pointer to the head of the list
}

// Add a node to the front of the list
func (ll *LinkedList) Add(value int) {
    newNode := &Node{Value: value}
    newNode.Next = ll.Head
    ll.Head = newNode
}

// Print the linked list
func (ll *LinkedList) Print() {
    current := ll.Head
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}

// Example of using a linked list
func linkedListExample() {
    ll := LinkedList{}
    ll.Add(1)
    ll.Add(2)
    ll.Print() // Output: 2 1
}

// ----------------------------------------
// 7. Queue
// ----------------------------------------

// A queue is a FIFO (First In, First Out) data structure.
type Queue []int

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(v int) {
    *q = append(*q, v)
}

// Dequeue removes and returns an element from the front of the queue
func (q *Queue) Dequeue() (int, bool) {
    if len(*q) == 0 {
        return 0, false
    }
    v := (*q)[0]
    *q = (*q)[1:]
    return v, true
}

// Example of using a queue
func queueExample() {
    q := Queue{}
    q.Enqueue(1)
    q.Enqueue(2)
    v, _ := q.Dequeue()
    fmt.Println(v) // Output: 1
}

// ----------------------------------------
// 8. Stack
// ----------------------------------------

// A stack is a LIFO (Last In, First Out) data structure.
type Stack []int

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    v := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return v, true
}

// Example of using a stack
func stackExample() {
    s := Stack{}
    s.Push(1)
    s.Push(2)
    v, _ := s.Pop()
    fmt.Println(v) // Output: 2
}

// ----------------------------------------
// 9. Binary Tree
// ----------------------------------------

// A binary tree is a tree data structure where each node has at most two children.
type TreeNode struct {
    Value int       // Value of the node
    Left  *TreeNode // Pointer to the left child
    Right *TreeNode // Pointer to the right child
}

// Insert adds a value to the tree
func (n *TreeNode) Insert(value int) {
    if value < n.Value {
        if n.Left == nil {
            n.Left = &TreeNode{Value: value}
        } else {
            n.Left.Insert(value)
        }
    } else {
        if n.Right == nil {
            n.Right = &TreeNode{Value: value}
        } else {
            n.Right.Insert(value)
        }
    }
}

// InOrderTraversal performs an in-order traversal of the tree
func (n *TreeNode) InOrderTraversal() {
    if n != nil {
        n.Left.InOrderTraversal()
        fmt.Print(n.Value, " ")
        n.Right.InOrderTraversal()
    }
}

// Example of using a binary tree
func binaryTreeExample() {
    root := &TreeNode{Value: 10}
    root.Insert(5)
    root.Insert(15)
    root.InOrderTraversal() // Output: 5 10 15
}

// ----------------------------------------
// 10. Graph
// ----------------------------------------

// A graph consists of nodes connected by edges.
type Graph struct {
    Nodes map[int][]int // Adjacency list representation
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 int) {
    g.Nodes[v1] = append(g.Nodes[v1], v2)
    g.Nodes[v2] = append(g.Nodes[v2], v1) // Undirected graph
}

// Example of using a graph
func graphExample() {
    g := Graph{Nodes: make(map[int][]int)}
    g.AddEdge(1, 2)
    g.AddEdge(1, 3)
    g.AddEdge(2, 4)
    fmt.Println(g.Nodes) // Output: map[1:[2 3] 2:[1 4] 3:[1] 4:[2]]
}

// Main function to run examples
func main() {
    arrayExample()
    sliceExample()
    mapExample()
    structExample()
    interfaceExample()
    linkedListExample()
    queueExample()
    stackExample()
    binaryTreeExample()
    graphExample()
}
