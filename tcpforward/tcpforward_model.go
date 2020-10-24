package tcpforward

import (
	"sync"
	"time"

	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
)

func (Self *Server) copyData(src, dst *gtcp.Conn, srcReader, dstReader ConnReadWrite, wg *sync.WaitGroup) {

	defer wg.Done()
	defer src.Close()
	defer dst.Close()
	for {
		b, err := srcReader.Read(src)
		if err != nil {
			glog.Error(err)
			break
		}
		_, err = dstReader.Write(dst, b)
		if err != nil {
			glog.Error(err)
			break
		}
	}
}

func (Self *Server) handleConn(src *gtcp.Conn) {

	Self.connectedCount++
	defer func() {
		Self.connectedCount--
	}()

	dst, err := gtcp.NewConn(Self.forwardAddr, time.Second*3)
	if err != nil {
		glog.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go Self.copyData(src, dst, Self.srcReadWrite, Self.dstReadWrite, &wg)
	go Self.copyData(dst, src, Self.dstReadWrite, Self.srcReadWrite, &wg)

	wg.Wait()
}

func (Self *Server) Serve() error {
	Self.tcpServer = gtcp.NewServer(Self.addr, Self.handleConn)
	return Self.tcpServer.Run()
}

func (Self *Server) ConnectedCount() int {
	return Self.connectedCount
}

func (Self *Server) Close() error {
	return Self.tcpServer.Close()
}
