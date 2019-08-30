package socketserver

import "net"

type SocketService interface {
	Serve(c net.Conn)
}
