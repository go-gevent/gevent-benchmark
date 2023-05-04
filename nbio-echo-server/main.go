package main

import (
	"flag"
	"fmt"

	"github.com/lesismal/nbio"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "5004", "server port")
	flag.Parse()

	g := nbio.NewGopher(nbio.Config{
		Network: "tcp",
		Addrs:   []string{":" + port},

		// Each poller hold a buffer with this size, poller num is not too many.
		// So, it's possible even if we set it to a large size.
		//ReadBufferSize: 1 * 1024 * 1024,

		MaxWriteBufferSize: 6 * 1024 * 1024,
	})
	// ajust kernel read/write buffer size
	g.OnOpen(func(c *nbio.Conn) {
		c.SetReadBuffer(1 * 1024 * 1024)
		c.SetWriteBuffer(1 * 1024 * 1024)
	})
	g.OnData(func(c *nbio.Conn, data []byte) {
		c.Write(data)
	})

	err := g.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}
	defer g.Stop()

	g.Wait()
}
