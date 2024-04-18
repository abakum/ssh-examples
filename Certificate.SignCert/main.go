package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Sign a certificate with a specific algorithm.
	privateKey, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		log.Fatal("unable to generate RSA key: ", err)
	}
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal("unable to get RSA public key: ", err)
	}
	caKey, err := rsa.GenerateKey(rand.Reader, 3072)
	if err != nil {
		log.Fatal("unable to generate CA key: ", err)
	}
	signer, err := ssh.NewSignerFromKey(caKey)
	if err != nil {
		log.Fatal("unable to generate signer from key: ", err)
	}
	mas, err := ssh.NewSignerWithAlgorithms(signer.(ssh.AlgorithmSigner), []string{ssh.KeyAlgoRSASHA256})
	if err != nil {
		log.Fatal("unable to create signer with algorithms: ", err)
	}
	certificate := ssh.Certificate{
		Key:      publicKey,
		CertType: ssh.UserCert,
	}
	if err := certificate.SignCert(rand.Reader, mas); err != nil {
		log.Fatal("unable to sign certificate: ", err)
	}
	// Save the public key to a file and check that rsa-sha-256 is used for
	// signing:
	// ssh-keygen -L -f <path to the file>
	fmt.Println(string(ssh.MarshalAuthorizedKey(&certificate)))
}
