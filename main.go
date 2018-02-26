package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		// exec itself
		cmd := exec.Command(os.Args[0], "--child")
		cmd.Start()
	} else {
		go ExitWait()
		for {
			h := time.Now().Hour()
			fmt.Println(h)
			if 5 >= h || h >= 21 {
				time.Sleep(9 * time.Minute)
				cmd := exec.Command("notify-send", "あと１分で画面がロックされます。")
				cmd.Start()
				time.Sleep(time.Minute)

				cmd = exec.Command("gnome-screensaver-command", "-l")
				cmd.Start()
			} else {
				time.Sleep(10 * time.Minute)
			}
		}
	}
}

func ExitWait() {
	ln, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	os.Exit(0)
}
