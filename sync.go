package main

// At the time of writing, I can't get the key-based connection to succeed, so
// deferring to Syncthing in the usage notes for now

// import (
// 	"log"
// 	"os"

// 	"github.com/pkg/sftp"
// 	"golang.org/x/crypto/ssh"
// )

// // syncNotes... syncs notes. Written with inspiration from the following two links:
// // https://pkg.go.dev/golang.org/x/crypto/ssh#AuthMethod
// // https://stackoverflow.com/a/63915669
// func syncNotes() {
// 	// A public key may be used to authenticate against the remote
// 	// server by using an unencrypted PEM-encoded private key file.
// 	//
// 	// If you have an encrypted private key, the crypto/x509 package
// 	// can be used to decrypt it.
// 	key, err := os.ReadFile("privateKeyPath")
// 	if err != nil {
// 		log.Fatalf("unable to read private key: %v", err)
// 	}

// 	// Create the Signer for this private key.
// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 		log.Fatalf("unable to parse private key: %v", err)
// 	}

// 	// var hostKey ssh.PublicKey
// 	config := &ssh.ClientConfig{
// 		User: "user",
// 		Auth: []ssh.AuthMethod{
// 			// Use the PublicKeys method for remote authentication.
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	// Connect to the remote server and perform the SSH handshake.
// 	client, err := ssh.Dial("tcp", "host:22", config)
// 	if err != nil {
// 		log.Fatalf("unable to connect: %v", err)
// 	}
// 	defer client.Close()

// 	// open an SFTP session over an existing ssh connection.
// 	sftp, err := sftp.NewClient(client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sftp.Close()

// 	// Open the source file
// 	srcFile, err := os.Open("/tmp/scrawl/a.md")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer srcFile.Close()

// 	// Create the destination file
// 	dstFile, err := sftp.Create("/tmp/a.md")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer dstFile.Close()
// }
