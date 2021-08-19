package mynet

type Config struct {
	host, protocol       string
	port, count, timeout int
}

func NewConfig(host string, port int, options ...Option) Config {
	c := Config{host: host, port: port, protocol: "tcp", count: -1, timeout: 100}
	for _, o := range options {
		o.apply(&c)
	}
	return c
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) Protocol() string {
	return c.protocol
}

func (c *Config) Port() int {
	return c.port
}

func (c *Config) CycleCount() int {
	return c.count
}

func (c *Config) DecCount() {
	c.count--
}

func (c *Config) Timeout() int {
	return c.timeout
}

type Option interface {
	apply(c *Config)
}

type protocolOption struct {
	protocol string
}

func (p protocolOption) apply(c *Config) {
	c.protocol = p.protocol
}

func WithProtocol(p string) Option {
	return protocolOption{protocol: p}
}

type hostOption struct {
	host string
}

func (h hostOption) apply(c *Config) {
	c.host = h.host
}

func WithHost(h string) Option {
	return hostOption{host: h}
}

type portOption struct {
	port int
}

func (p portOption) apply(c *Config) {
	c.port = p.port
}

func WithPort(p int) Option {
	return portOption{port: p}
}

type cycleOption struct {
	count int
}

func (co cycleOption) apply(c *Config) {
	c.count = co.count
}

func WithCycleCount(c int) Option {
	return cycleOption{count: c}
}

type timeoutOption struct {
	timeout int
}

func (t timeoutOption) apply(c *Config) {
	if t.timeout > 0 {
		c.timeout = t.timeout
	}
}

func WithTimeout(t int) Option {
	return timeoutOption{timeout: t}
}
