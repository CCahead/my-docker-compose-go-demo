package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
  

   

    
    // 连接到服务端
    conn, err := net.Dial("tcp", "server:5050")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    for {
    // 从命令行读取输入
    fmt.Print("Enter text: ")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    // 发送到服务端
    fmt.Fprintf(conn, text+"\n")

    // 读取服务端的响应
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: " + message)
}
}