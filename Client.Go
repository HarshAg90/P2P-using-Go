package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var ip string
var port string
var clientsList map[string]string

func main() {
	// Get the client's public IP address by making an HTTP request to an external service
	// go func() {
	func() {
		resp, err := net.Dial("tcp", "ifconfig.me:80")
		if err != nil {
			fmt.Println("Could not determine public IP:", err)
			return
		}
		defer resp.Close()
		fmt.Fprintf(resp, "GET /ip HTTP/1.0\r\nHost: ifconfig.me\r\n\r\n")
		scanner := bufio.NewScanner(resp)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				if scanner.Scan() {
					ip = scanner.Text()
				}
				break
			}
		}
		if ip == "" {
			fmt.Println("error getting ip addr")
			return
		}
		// if ip != "" {
		// 	fmt.Println("Your public IP address is:", ip)
		// }
	}()
	// Wait for the IP to be set
	fmt.Printf("Enter desired port, '5000+' recomnded:")
	portt := bufio.NewScanner(os.Stdin)
	portt.Scan()
	port = portt.Text()
	_, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Invalid port number. Exiting.")
		return
	}

	// creating a TCP connection to the main server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the chat server.")

	// async Read loop from server
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			idPost := strings.Split(scanner.Text(), `|`)
			if idPost[0] == "metadata" {
				// fmt.Println("Received IP:", idPost[1])
				clientsList = ParseClientsList(idPost[1])
				// if clientsList != nil {
				// 	// ip = jsonMap["ip"][0].(string)
				// 	for i, v := range clientsList {
				// 		fmt.Printf("Peer %s IP: %s\n", i, v)
				// 	}
				// } else {
				// 	go startLocalChatListener()
				// 	fmt.Println("No IP addresses found in metadata, starting local chat listener.")
				// }
			} else {
				fmt.Println(scanner.Text())
			}
		}
		fmt.Println("Disconnected from server.")
		os.Exit(0) // exit if connection drops
	}()

	// Read from user input
	fmt.Printf("enter your name:")
	input := bufio.NewScanner(os.Stdin)
	flag := false
	for input.Scan() {
		if input.Text() == "exit" {
			fmt.Println("Exiting chat...")
			break
		}
		if input.Text() == "/peers" {
			printPeerList()
			continue
		}
		if !flag {
			_, err := fmt.Fprintln(conn, input.Text()+"|"+ip+":"+port)
			flag = true
			if err != nil {
				fmt.Println("Error sending message:", err)
				break
			}
		} else {
			_, err := fmt.Fprintln(conn, input.Text())
			if err != nil {
				fmt.Println("Error sending message:", err)
				break
			}
		}
	}

	fmt.Println("Exiting chat...")
}

func printPeerList() {
	if clientsList != nil {
		// ip = jsonMap["ip"][0].(string)
		fmt.Printf("\nPeer List\n")
		ind := 1
		for i, v := range clientsList {
			fmt.Printf("#%d - %s IP: %s\n", ind, i, v)
			ind++
		}
	} else {
		go startLocalChatListener()
		fmt.Println("No IP addresses found in metadata, starting local chat listener.")
	}
}

func startLocalChatListener() {
	ln, err := net.Listen("tcp", ":"+port) // Listen on all interfaces, port 9000
	if err != nil {
		fmt.Println("Error starting local chat listener:", err)
		return
	}
	fmt.Println("Local chat listener started on port 9000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting local chat connection:", err)
			continue
		}
		go handleLocalChat(conn)
	}
}

func handleLocalChat(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted local chat connection from", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("[Local Chat]", scanner.Text())
	}
}

func connectToLocalChat(ip string) {
	conn, err := net.Dial("tcp", ip+":9000")
	if err != nil {
		fmt.Println("Error connecting to local chat:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to local chat at", ip+":9000")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		_, err := fmt.Fprintln(conn, input.Text())
		if err != nil {
			fmt.Println("Error sending local chat message:", err)
			break
		}
	}
}

func NormalizeIP(ip string) string {
	// Remove brackets if present
	ip = strings.Trim(ip, "[]")
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return ip // Return as-is if not a valid IP
	}
	// If IPv4, return in dotted format
	if ipv4 := parsed.To4(); ipv4 != nil {
		return ipv4.String()
	}
	// If IPv6 loopback, return IPv4 loopback
	if parsed.IsLoopback() {
		return "127.0.0.1"
	}
	// Otherwise, return the string form
	return parsed.String()
}

func ParseClientsList(msg string) map[string]string {
	clients := make(map[string]string)
	entries := strings.Split(msg, ",")
	for _, entry := range entries {
		parts := strings.SplitN(entry, "-", 2)
		if len(parts) == 2 {
			// Split address and port
			addrParts := strings.Split(parts[0], ":")
			if len(addrParts) >= 2 {
				ipPart := strings.Join(addrParts[:len(addrParts)-1], ":")
				portPart := addrParts[len(addrParts)-1]
				ipPart = NormalizeIP(ipPart)
				clients[fmt.Sprintf("%s:%s", ipPart, portPart)] = parts[1]
			} else {
				clients[parts[0]] = parts[1]
			}
		}
	}
	return clients
}

// request TCP connection to server with our IP and our name
// create socket connection to server
// once list of IP addresses is received, Verify witch one to connect to by name (also given by server)
// dirrectly connect to the IP address of the server
// UDP Hole punching into another client's IP address,
// listn for incoming UDP messages from other clients
// send UDP messages to other clents
