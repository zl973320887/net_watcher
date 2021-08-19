package main

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"watcher/mynet"
	"watcher/mynet/tcp"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout int    `yaml:"timeout"`
	Count   int    `yaml:"count"`
	Echo    bool   `yaml:"echo"`
}

var (
	host, protocol string
	port, count    int
	timeout        int
	echo           bool
	cfg            Config
)

func init() {
	flag.StringVar(&host, "host", "", "destination host")
	flag.IntVar(&port, "port", 0, "destination port")
	flag.StringVar(&protocol, "protocol", "tcp", "network protocol")
	flag.IntVar(&count, "count", -1, "test count")
	flag.IntVar(&timeout, "timeout", -1, "timeout standard")
	flag.BoolVar(&echo, "echo", true, "echo switch")
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
		echo = cfg.Echo
		log.Printf("config: [[[ host:%s, port:%d, count:%d, timeout:%d, echo:%v", host, port, count, timeout, echo)
	}
	w := mynet.NewWatcher(tcp.New(mynet.NewConfig(host, port, mynet.WithTimeout(timeout), mynet.WithCycleCount(count), mynet.WithEcho(echo))))
	w.Watch()
}
