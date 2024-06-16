package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func GetResponse(conn net.Conn, done chan<- bool) {
	messageScanner := bufio.NewScanner(conn)
	for messageScanner.Scan() {
		fmt.Println(messageScanner.Text())
	}

	if err := messageScanner.Err(); err != nil {
		fmt.Printf("Server disconnected: %v\n", err)
		done <- true
		return
	}
}

func ReadInput(reader *bufio.Reader, inputChan chan<- string) {
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		inputChan <- message
	}
}

func main() {
	conn, err := net.Dial("tcp", ":7007")
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Name: ")
	name, _ := reader.ReadString('\n')
	_, err = fmt.Fprintf(conn, "%s", name)
	if err != nil {
		fmt.Printf("Failed to send name: %v\n", err)
		return
	}

	done := make(chan bool)
	inputChan := make(chan string)
	go GetResponse(conn, done)
	go ReadInput(reader, inputChan)

	for {
		select {
		case <-done:
			fmt.Println("Server disconnected, exiting...")
			return
		case message := <-inputChan:
			if strings.ToLower(message) == "quit\n" || strings.ToLower(message) == "exit\n" {
				fmt.Printf("%s signing off :)\n", name)
				conn.Close()
				return
			}

			_, err = fmt.Fprintf(conn, message)
			if err != nil {
				fmt.Printf("Failed to send message: %v\n", err)
				return
			}
		}
	}
}
