package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	USER string = ""
	PASS string = ""
	PORT int    = 22
	HOST string = ""
)

func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 30 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

// func main() {
// 	session, err := connect(USER, PASS, HOST, PORT)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer session.Close()

// 	// session.Stdout = os.Stdout
// 	// session.Stderr = os.Stderr
// 	// session.Run("ls /")

// 	out, err := session.CombinedOutput("ls /")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(out))
// }
