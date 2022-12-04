package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"time"
)

// Terminal 终端结构体
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
		log.Fatalf("error connect: %v", err)
	}

	terminal.Client = client
	log.Println("connect successful.")
}

// NewSession 开启新的会话
func (terminal *Terminal) NewSession() {
	session, err := terminal.Client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("session error for stdin: %v", err)
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("session error for stdout: %v", err)
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("session session for stderr: %v", err)
	}
	go io.Copy(os.Stderr, stderr)

	modes := ssh.TerminalModes{
		// 禁用回显
		ssh.ECHO: 0,
		// 输入速度
		ssh.TTY_OP_ISPEED: 14400,
		// 输出速度
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err = session.RequestPty("xterm", 45, 140, modes); err != nil {
		log.Fatalf("session error for request pty: %v", err)
	}

	if err = session.Shell(); err != nil {
		log.Fatalf("session error for shell command: %v", err)
	}

	if err = session.Wait(); err != nil {
		log.Fatalf("session error for wait: %v", err)
	}
}
