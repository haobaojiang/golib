package tcpforward

import "github.com/gogf/gf/net/gtcp"

type ConnReadWrite interface {
	Read(conn *gtcp.Conn) ([]byte, error)
	Write(conn *gtcp.Conn, data []byte) (int, error)
}

type Server struct {
	tcpServer    *gtcp.Server
	addr         string
	forwardAddr  string
	srcReadWrite ConnReadWrite
	dstReadWrite ConnReadWrite

	//
	connectedCount int
}

func New(addr string, forwardAddr string, srcReadwrite ConnReadWrite, dstReadWrite ConnReadWrite) *Server {
	return &Server{
		addr:         addr,
		forwardAddr:  forwardAddr,
		srcReadWrite: srcReadwrite,
		dstReadWrite: dstReadWrite,
	}
}
