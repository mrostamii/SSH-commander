package main

import (
	"fmt"
	"time"

	"github.com/appleboy/easyssh-proxy"
)

func main() {
	ssh := &easyssh.MakeConfig{
		User:     "root",
		Server:   "192.168.1.35",
		Password: "11111",
		Port:     "22",
		Timeout:  60 * time.Second,
	}

	stdout, stderr, done, err := ssh.Run("shutdown now", 60*time.Second)

	if err != nil {
		panic("ERR: (Can't run remote command): " + err.Error())
	} else {
		fmt.Println("	LOG (Done):", done, "\n	LOG (stdout):", stdout, "\n	LOG (stderr):", stderr)
	}
}
