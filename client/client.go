package client

import (
	"net"
)

const BUF_SIZE = 1024

type KeyValueClient struct {
	addr string
	conn net.Conn
}

func New(addr string) KeyValueClient {
	return KeyValueClient{addr, nil}
}

func (r *KeyValueClient) Connect() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", r.addr)

	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		return err
	}

	r.conn = conn

	return nil
}

func (r *KeyValueClient) Set(key, value string) (string, error) {
	r.conn.Write([]byte("set " + key + " " + value))

	buf := make([]byte, BUF_SIZE)

	resLen, err := r.conn.Read(buf)

	if err != nil {
		return "", err
	}

	return string(buf[0:resLen]), nil
}

func (r *KeyValueClient) Get(key string) (string, error) {
	r.conn.Write([]byte("get " + key))

	buf := make([]byte, BUF_SIZE)

	resLen, err := r.conn.Read(buf)

	if err != nil {
		return "", err
	}

	return string(buf[0:resLen]), nil
}

