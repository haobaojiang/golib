package gtcpforward

import (
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"sync"
	"time"
)

func(Self* Server) handleForward(conn *gtcp.Conn, target *gtcp.Conn, wg *sync.WaitGroup) {

	defer wg.Done()
	defer target.Close()
	defer conn.Close()
	for {
		b, err := conn.Recv(0)
		if err != nil {
			glog.Error(err)
			break
		}
		encrypted, err := Self.enc.Encrypt(b)
		if err != nil {
			glog.Error(err)
			break
		}
		err = target.SendPkg(encrypted)
		if err != nil {
			glog.Error(err)
			break
		}
	}
}

func(Self* Server) handleBackWard(conn *gtcp.Conn, target *gtcp.Conn, wg *sync.WaitGroup) {

	defer wg.Done()
	defer target.Close()
	defer conn.Close()
	for {
		// read from target
		b, err := target.RecvPkg()
		if err != nil {
			glog.Error(err)
			break
		}

		// decrypt data
		decrypted, err := Self.dec.Decrypt(b)
		if err != nil {
			glog.Error(err)
			break
		}
		_, err = conn.Write(decrypted)
		if err != nil {
			glog.Error(err)
			break
		}
	}
}

func (Self *Server) handleConn(client *gtcp.Conn) {

	loginSvcConn, err := gtcp.NewConn(Self.forwardAddr, time.Second*3)
	if err != nil {
		glog.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go Self.handleForward(loginSvcConn, client, &wg)
	go Self.handleBackWard(loginSvcConn, client, &wg)

	wg.Wait()
}

func (Self *Server) Serve() error {
	return gtcp.NewServer(Self.addr, Self.handleConn).Run()
}

