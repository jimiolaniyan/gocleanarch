package socketserver

import (
	"fmt"
	"net"
	"strconv"
)

type SocketServer struct {
	port     int
	service  SocketService
	running  bool
	listener net.Listener
}

func NewSocketServer(port int, service SocketService) (*SocketServer, error) {
	var err error
	l, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	if err != nil {
		return nil, err
	}

	return &SocketServer{port: port, service: service, listener: l}, nil
}

func (ss *SocketServer) Port() int {
	return ss.port
}

func (ss *SocketServer) Service() SocketService {
	return ss.service
}

func (ss *SocketServer) start() {
	go acceptConnections(ss)
	ss.running = true
}

func acceptConnections(server *SocketServer) {
	for server.running {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go server.service.serve(conn)
	}
}

func (ss *SocketServer) Running() bool {
	return ss.running
}

func (ss *SocketServer) stop() {
	var err error

	if ss.running {
		err = ss.listener.Close()
	}

	if err != nil {
		fmt.Println(err)
	}

	ss.running = false
}
