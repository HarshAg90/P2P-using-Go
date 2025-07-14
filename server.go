package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	clients     = make(map[net.Conn]string)
	broadcast   = make(chan string)
	clientsList = make(map[string]string)
	mutex       sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go handleBroadcast()

	fmt.Println("Server started on :8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// bound to run after this function ends fully or early by error
	defer conn.Close()

	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	ip_addr := strings.Split(name, "|")[1] // split to get name only if port is included
	name = strings.Split(name, "|")[0]     // split to get name only if port is included

	// lets only 1 thread run, stops other threads from writing to the map at the same time, helps to prevent race conditions and data connection for multiple threads working on same data
	mutex.Lock()
	clients[conn] = name
	clientsList[ip_addr] = name
	mutex.Unlock()

	broadcast <- fmt.Sprintf("Welcome %s.\nType '/peers' to reveal all connected peer\nType '/connected 1' to connect to #1 peer\n", name)
	// Print the current clients map to the server console with key-value pairs
	mutex.Lock()
	fmt.Println("Current clients (conn -> name):")
	for conn, clientName := range clients {
		fmt.Printf("%v -> %s\n", conn.RemoteAddr(), clientName)
	}
	mutex.Unlock()

	for {
		// Send the updated clients list to all clients
		// metadata|12343=harsh,234532=harsh2
		if len(clientsList) != 0 {
			clientList := "metadata|"
			mutex.Lock()
			clientli := []string{}
			for ip, clientName := range clientsList {
				// str := ip.RemoteAddr().String() + "-" + clientName
				str := ip + "-" + clientName
				clientli = append(clientli, str)
			}
			mutex.Unlock()
			clientList += strings.Join(clientli, ",")
			broadcast <- clientList
		}

		// for regular chat between clients
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msg = strings.TrimSpace(msg)
		broadcast <- fmt.Sprintf("[%s]: %s", name, msg)

	}

	mutex.Lock()
	delete(clients, conn)
	delete(clientsList, ip_addr)
	mutex.Unlock()
	broadcast <- fmt.Sprintf("%s left the chat", name)
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mutex.Unlock()
	}
}

// create TCP endpoint to connect to server
// create a socket connection to incoming connections and store their iP in list,
// one a new connection connects to server about their IP,  send a message to all connected clients
