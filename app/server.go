package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func handleConnection(conn net.Conn) {
  defer conn.Close()

  reader := bufio.NewReader(conn)
  message, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Error reading: ", err)
    return
  }

  fmt.Print("Message received: ", string(message))
  header := strings.Split(message, " ")

  fmt.Println(header[1])

  response := ""
  if header[1] == "/" {
    response = "HTTP/1.1 200 OK\r\n\r\n"
  } else {
    response = "HTTP/1.1 404 Not Found\r\n\r\n"
  }

  _, err = conn.Write([]byte(response))
  if err != nil {
    fmt.Println("Error writing in connection: ", err.Error())
  }
}


func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
	  fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
  }
  defer l.Close()
	
  for {
    conn, err := l.Accept()
	  if err != nil {
	 	  fmt.Println("Error accepting connection: ", err.Error())
      continue
	  }

    go handleConnection(conn)
  }

}
