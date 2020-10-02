package tcpforward

import "github.com/gogf/gf/net/gtcp"

type ConnReadWrite interface {
	Read(conn *gtcp.Conn) ([]byte, error)
	Write(conn *gtcp.Conn, data []byte) (int, error)
}

type Server struct {
	addr         string
	forwardAddr  string
	srcReadWrite ConnReadWrite
	dstReadWrite ConnReadWrite
}

func New(addr string, forwardAddr string, srcReadwrite ConnReadWrite, dstReadWrite ConnReadWrite) *Server {
	return &Server{
		addr:         addr,
		forwardAddr:  forwardAddr,
		srcReadWrite: srcReadwrite,
		dstReadWrite: dstReadWrite,
	}
}
