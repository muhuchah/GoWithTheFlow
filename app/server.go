package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func handleContentType(filePath string) string {
  contentType := "text/plain"
  if strings.HasSuffix(filePath, ".html") {
    contentType = "text/html"
  } else if strings.HasSuffix(filePath, ".css") {
    contentType = "text/css"
  } else if strings.HasSuffix(filePath, ".js") {
    contentType = "application/javascript"
  }

  return contentType
}


func handleConnection(conn net.Conn) {
  defer conn.Close()

  reader := bufio.NewReader(conn)
  requestLine, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Error reading: ", err)
    return
  }

  fmt.Print("Message received: ", string(requestLine))
  headerParts := strings.Split(requestLine, " ")
  if len(headerParts) < 2 {
    fmt.Println("Invalid request")
    return
  }
  
  method := headerParts[0]
  path := headerParts[1]

  if method == "GET" {
    filePath := "static" + path
    if path == "/" {
      filePath = "static/index.html"
    }

    file, err := os.ReadFile(filePath)
    if err != nil {
      // File not found
      response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\n404 Not Found"
      conn.Write([]byte(response))
      return
    }

    ContentType := handleContentType(filePath)

    response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s", ContentType, len(file), file)
    conn.Write([]byte(response))
  } else {
    response := "HTTP/1.1 405 Method Not Allowed\r\nContent-Type: text/plain\r\n\r\n405 Method Not Allowed"
		conn.Write([]byte(response))
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

