package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func GetResponse(conn net.Conn) {
	messageScanner := bufio.NewScanner(conn)
	messageScanner.Scan()

	fmt.Println(messageScanner.Text())
}

func main() {
	conn, err := net.Dial("tcp", ":7007")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Name: ")
	name, _ := reader.ReadString('\n')
	_, err = fmt.Fprintf(conn, "%s", name)
	if err != nil {
		panic(err)
	}
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
			conn.Close()
		}

		if strings.ToLower(message) == "quit\n" || strings.ToLower(message) == "exit\n" {
			fmt.Printf("%s signing off :)\n", name)
			break
		}

		_, err = fmt.Fprintf(conn, message)
		if err != nil {
			panic(err)
		}

		GetResponse(conn)
	}

}
