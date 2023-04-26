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
		Network:            "tcp",
		Addrs:              []string{":" + port},
		MaxWriteBufferSize: 6 * 1024 * 1024,
	})

	g.OnData(func(c *nbio.Conn, data []byte) {
		_, err := c.Write(append([]byte{}, data...))
		if err != nil {
			return
		}
	})

	err := g.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}
	defer g.Stop()

	g.Wait()
}
