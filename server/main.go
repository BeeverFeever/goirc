package main

import (
    "bufio"
    "fmt"
    "io"
    "net"
    "os"
)

func handleconnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("connection found")
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err == io.EOF {
            fmt.Println("Connection closed")
            return
        }
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error: ", err)
            return
        }
        fmt.Println(msg)
        conn.Write([]byte(msg))
    }
}

func main() {
    ln, err := net.Listen("tcp", "localhost:6667")
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error: ", err)
        return
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error: ", err)
            return
        }
        go handleconnection(conn)
    }
}
