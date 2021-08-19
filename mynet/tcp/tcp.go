package tcp

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"net"
	"time"
	"watcher/mynet"
)

type Tcp struct {
	c mynet.Config
}

func New(conf mynet.Config) *Tcp {
	return &Tcp{c: conf}
}

func (n *Tcp) Connect() (error, time.Duration) {
	address := fmt.Sprintf("%s:%d", n.c.Host(), n.c.Port())
	nets := new(net.Dialer)
	nets.Timeout = time.Millisecond * 3000
	startTime := time.Now()
	conn, err := nets.Dial(n.c.Protocol(), address)
	if err != nil {
		return err, 1000000
	}
	conn.Write([]byte("hello I'm test!"))
	duration := time.Since(startTime).Round(time.Microsecond)
	defer conn.Close()
	return nil, duration
}

func (n *Tcp) Ping() {
	forever := true
	if n.c.CycleCount() > 0 {
		forever = false
	}
	for forever || n.c.CycleCount() > 0 {
		date := time.Now().Format("2006-01-02 15:04:05")
		err, timeout := n.Connect()
		if err != nil {
			if n.c.Echo() {
				fmt.Printf("%s Connect to %s:%s %s", date, color.Blue(n.c.Host()), color.Blue(n.c.Port()), color.Red("Connection timed out\n"))
			} else {
				log.Printf("%s Connect to %s:%d %s", date, n.c.Host(), n.c.Port(), "Connection timed out\n")
			}
		} else {
			if n.c.Echo() {
				fmt.Printf("%s Connected to  %s:%s protocol=%s timeout=%s\n",
					date, color.Green(n.c.Host()), color.Green(n.c.Port()), color.Green(n.c.Protocol()), color.Green(timeout.String()))
			}
		}
		if timeout.Milliseconds() >= int64(n.c.Timeout()) {
			log.Printf("%s Connected to  %s:%d protocol=%s timeout=%s\n",
				date, n.c.Host(), n.c.Port(), n.c.Protocol(), timeout.String())
		}
		time.Sleep(1000 * time.Millisecond)
		if !forever {
			n.c.DecCount()
		}
	}
}
