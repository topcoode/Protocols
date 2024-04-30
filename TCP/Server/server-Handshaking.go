package main

import (
    "fmt"
    "net"
)

func main() {
    // Start TCP server
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer ln.Close()
    fmt.Println("Server listening on port 8080")

    // Accept incoming connections
    conn, err := ln.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err.Error())
        return
    }
    defer conn.Close()
    fmt.Println("Client connected")

    // Read data from client
    buf := make([]byte, 1024)
    _, err = conn.Read(buf)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
        return
    }
    fmt.Println("Received:", string(buf))

    // Send response to client
    conn.Write([]byte("Message received"))
    fmt.Println("Response sent")
}

