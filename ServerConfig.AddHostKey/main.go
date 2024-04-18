package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Minimal ServerConfig supporting only password authentication.
	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in
			// a production setting.
			if c.User() == "testuser" && string(pass) == "tiger" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes, err := os.ReadFile("id_rsa")
	if err != nil {
		log.Fatal("Failed to load private key: ", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key: ", err)
	}
	// Restrict host key algorithms to disable ssh-rsa.
	signer, err := ssh.NewSignerWithAlgorithms(private.(ssh.AlgorithmSigner), []string{ssh.KeyAlgoRSASHA256, ssh.KeyAlgoRSASHA512})
	if err != nil {
		log.Fatal("Failed to create private key with restricted algorithms: ", err)
	}
	config.AddHostKey(signer)
}
