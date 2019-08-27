package socketserver

import "net"

type SocketService interface {
	serve(c net.Conn)
}
