package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
)

func main() {
	// 1. Listen on a UDP port
	addr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, _ := net.ListenUDP("udp", &addr)
	defer conn.Close()

	fmt.Println("Server listening on UDP :12345...")

	buf := make([]byte, 1024)
	for {
		// 2. Read packet
		n, _, _ := conn.ReadFromUDP(buf)
		data := string(buf[:n])

		// 3. Parse "x,y,pressure"
		parts := strings.Split(data, ",")
		if len(parts) >= 2 {
			x, _ := strconv.ParseFloat(parts[0], 64)
			y, _ := strconv.ParseFloat(parts[1], 64)

			// 4. Move the mouse
			// Note: You may need to scale these coordinates
			// based on your PC screen resolution
			robotgo.Move(int(x), int(y))
		}
	}
}
