package key_value_client

import (
	"net"
)

const BUF_SIZE = 1024

type KyeValueClient struct {
	addr string
	conn net.Conn
}

func New(addr string) KyeValueClient {
	return KyeValueClient{addr, nil}
}

func (r *KyeValueClient) Connect() error {
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

func (r *KyeValueClient) Set(key, value string) (string, error) {
	r.conn.Write([]byte("set " + key + " " + value))

	buf := make([]byte, BUF_SIZE)

	resLen, err := r.conn.Read(buf)

	if err != nil {
		return "", err
	}

	return string(buf[0:resLen]), nil
}

func (r *KyeValueClient) Get(key string) (string, error) {
	r.conn.Write([]byte("get " + key))

	buf := make([]byte, BUF_SIZE)

	resLen, err := r.conn.Read(buf)

	if err != nil {
		return "", err
	}

	return string(buf[0:resLen]), nil
}

