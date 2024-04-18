package main

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := "testuser"
	NumberOfPrompts := 3

	// Normally this would be a callback that prompts the user to answer the
	// provided questions
	Cb := func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
		return []string{"answer1", "answer2"}, nil
	}

	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            user,
		Auth: []ssh.AuthMethod{
			ssh.RetryableAuthMethod(ssh.KeyboardInteractiveChallenge(Cb), NumberOfPrompts),
		},
	}

	host := "mysshserver"
	netConn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	sshConn, _, _, err := ssh.NewClientConn(netConn, host, config)
	if err != nil {
		log.Fatal(err)
	}
	_ = sshConn
}
