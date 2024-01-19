package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // 在localhost的8080端口上启动TCP监听
    ln, err := net.Listen("tcp", ":5050")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // 延迟关闭监听器
    defer ln.Close()

    fmt.Println("Listening on :5050")
    for {
        // 接受一个客户端连接
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // 处理连接
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    // 读取客户端发来的消息
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }
    fmt.Print("Message received:", string(message))
    // 发送一个响应消息到客户端
    conn.Write([]byte("Message received.\n"))
    // 关闭连接
    // conn.Close()
}
