package pool

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}
func startNetworkDaemon2() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot lisen: %v", err)
		}
		server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintf(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}
