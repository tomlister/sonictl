package main

import (
	"flag"
	"net"
)

func e(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	up := flag.Bool("up", false, "up")
	down := flag.Bool("down", false, "down")
	flag.Parse()
	if (*up == true && *down == false) {
		c, err := net.Dial("unix", "/tmp/sonic.sock")
		e(err)
		defer c.Close()
		_, err = c.Write([]byte("up\n"))
		e(err)
	} else if (*up == false && *down == true) {
		c, err := net.Dial("unix", "/tmp/sonic.sock")
		e(err)
		defer c.Close()
		_, err = c.Write([]byte("down\n"))
		e(err)
	}
}