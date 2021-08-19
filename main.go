package main

import (
	"flag"
	"fmt"
	"gaping/mynet"
	"gaping/mynet/tcp"
	"github.com/labstack/gommon/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout int    `yaml:"timeout"`
	Count   int    `yaml:"count"`
}

var (
	host, protocol string
	port, count    int
	timeout        int
	cfg            Config
)

func init() {
	flag.StringVar(&host, "h", "", "Dest host ipaddress")
	flag.IntVar(&port, "p", 0, "Dest host port")
	flag.StringVar(&protocol, "protocol", "tcp", "Set network protocol")
	flag.IntVar(&count, "c", 100000000, "Tcp test conuts")
}

func main() {
	flag.Parse()
	if host == "" || port == 0 {
		fmt.Printf(color.Red("host or port is empty. will try reading config from file\n"))
		content, err := ioutil.ReadFile("./etc/app.yaml")
		if err != nil {
			log.Printf("read config content error:%s", err.Error())
			return
		}
		err = yaml.Unmarshal(content, &cfg)
		if err != nil {
			log.Printf("parse config error:%s", err.Error())
			return
		}
		host = cfg.Host
		port = cfg.Port
		count = cfg.Count
		timeout = cfg.Timeout
	}
	w := mynet.NewWatcher(tcp.New(mynet.NewConfig(host, port, mynet.WithTimeout(timeout), mynet.WithCycleCount(count))))
	w.Watch()
}
