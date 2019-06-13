package main

import (
    "fmt"
    "net"
    "os"
    "bytes"
)

const (
    CONN_HOST = "114.32.187.6"
    CONN_PORT = "8085"
    CONN_TYPE = "tcp"
)

func main() {
    listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        printLog("Error listening:", err.Error())
        os.Exit(1)
    }

    defer listener.Close()
    printLog("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        conn, err := listener.Accept()
        if err != nil {
            printLog("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
	printLog(conn.RemoteAddr().String()+" is connected")
	for{
		test, err := read(conn)
		if err != nil {
			printLog(conn.RemoteAddr().String()+" is disconnected: ", err.Error())
			defer conn.Close()
			break
		}
		printLog(test)
	}
}

var logSn = 0
func printLog(format string, args ...interface{}) {
	fmt.Printf("%d: %s \n", logSn, fmt.Sprintf(format, args...))
	logSn++
}

func read(conn net.Conn) (string, error) { 
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			defer conn.Close()
			return "", err
		}

		readByte := readBytes[0]
		if readByte == ' ' {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error){
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(' ')
	return conn.Write(buffer.Bytes())
}