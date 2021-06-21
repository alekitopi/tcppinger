package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var help_msg string = `Uso: ./pinger <ip> <puerto> [<opciones>]

Opciones:
-h    Muestra el mensaje de ayuda.
-t    Modifica el delay entre cada pingeo.

`

var (
	host    string
	port    string
	timeout string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	raw_args := os.Args
	if len(raw_args) < 3 {
		fmt.Printf(help_msg)
		os.Exit(1)
	} else {
		host = raw_args[1]
		port = raw_args[2]
	}
	// set defaults
	timeout = "3"
	for i, a := range raw_args[1:] {
		if !strings.HasPrefix(a, "-") {
			continue
		} else if a == "-h" {
			fmt.Println(help_msg)
			os.Exit(0)
		} else if a == "-t" {
			timeout = raw_args[i+2]
		} else {
			fmt.Println(help_msg)
			fmt.Printf("argumento no válido: %v\n", a)
			os.Exit(1)
		}
	}
	t, err := strconv.Atoi(timeout)
	check(err)
	target := host + ":" + port

	for {
		start_time := time.Now()
		conn, err := net.DialTimeout("tcp", target, time.Duration(t)*time.Second)

		end_time := time.Now()
		elapsed_time := float64(end_time.Sub(start_time)) / float64(time.Millisecond)

		if err != nil {
			fmt.Printf("[!] %v (%v/tcp) - no disponible - %f ms\n", host, port, elapsed_time)
		} else {
			conn.Close()
			fmt.Printf("[!] %v (%v/tcp) - disponible - %f ms\n", host, port, elapsed_time)
		}

		time.Sleep(time.Duration(t) * time.Second)
	}
	os.Exit(0)
}