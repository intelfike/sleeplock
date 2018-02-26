package main

import (
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
		for {
			h := time.Now().Hour()
			if 5 >= h || h >= 22 {
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
