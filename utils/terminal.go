package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

type Terminal struct {
	Name   string
	Host   string
	Port   string
	User   string
	Pass   string
	Client *ssh.Client
}

// Connect 建立连接
func (terminal *Terminal) Connect() {
	var (
		client *ssh.Client
		err    error
	)
	config := ssh.ClientConfig{
		User:            terminal.User,
		Auth:            []ssh.AuthMethod{ssh.Password(terminal.Pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%s", terminal.Host, terminal.Port)

	if client, err = ssh.Dial("tcp", addr, &config); err != nil {
		log.Fatalf("error connect:%v\n", err)
	}

	terminal.Client = client
	log.Println("connect successful.")
}

// NewSession 开启新的会话
func (terminal *Terminal) NewSession() {
	log.Println("starting new session.")
	session, err := terminal.Client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s\n", err.Error())
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("session request pty error: %s\n", err.Error())
	}

	if err = session.Shell(); err != nil {
		log.Fatalf("session shell error: %s\n", err.Error())
	}

	if err = session.Wait(); err != nil {
		log.Fatalf("session error: %s\n", err.Error())
	}
}
