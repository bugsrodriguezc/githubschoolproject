package main

import "fmt"

func main() {
    // Define a struct to represent a node in a linked list
    type Node struct {
        Value int
        Next  *Node
    }

    // Create a new linked list
    var head *Node = &Node{1, nil}

    // Traverse the linked list and print each element
    for next := head; next != nil; next = next.Next {
        fmt.Println("Value:", next.Value)
    }
}
