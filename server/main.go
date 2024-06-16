package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	name string
}

type Message struct {
	name    string
	content string
}

func Server(conn net.Conn, clients chan<- Client, messages chan<- Message) {

	nameScanner := bufio.NewScanner(conn)
	nameScanner.Scan()
	clientName := nameScanner.Text()

	clients <- Client{conn: conn, name: clientName}

	for {
		buffer := make([]byte, 1024) // Adjust buffer size as needed
		n, err := conn.Read(buffer)
		if err == net.ErrWriteToConnected {
			panic(err)
			// No data available yet, continue checking for other clients
		} else if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}

		// Process received message
		content := string(buffer[:n])
		if strings.ToLower(content) == "exit\n" {
			break
		}
		messages <- Message{name: clientName, content: content}
		//fmt.Fprintf(conn, toWrite)
	}
}

func Broadcast(clientList map[string]Client, msg Message) {
	for _, c := range clientList {
		if msg.name == c.name {
			continue
		}

		_, err := c.conn.Write([]byte(fmt.Sprintf("%v: %v", msg.name, msg.content)))
		if err != nil {
			panic(err)
		}
	}

}

func main() {

	listener, err := net.Listen("tcp", ":7007")

	if err != nil {
		fmt.Printf("Error: ", err)
	}

	defer listener.Close()

	clients := make(chan Client)
	messages := make(chan Message)

	ClientList := make(map[string]Client)
	go func() {
		for msg := range messages {
			Broadcast(ClientList, msg)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error: %v", err.Error())
		}
		go Server(conn, clients, messages)

		newClient := <-clients
		ClientList[newClient.name] = newClient
	}
}
