package main

import (
    "fmt"
    "net"
)

func main() {
    // Connect to server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        return
    }
    defer conn.Close()

    // Send data to server
    message := []byte("Hello, server")
    _, err = conn.Write(message)
    if err != nil {
        fmt.Println("Error sending:", err.Error())
        return
    }
    fmt.Println("Message sent")

    // Read response from server
    buf := make([]byte, 1024)
    _, err = conn.Read(buf)
    if err != nil {
        fmt.Println("Error reading response:", err.Error())
        return
    }
    fmt.Println("Response from server:", string(buf))
}

