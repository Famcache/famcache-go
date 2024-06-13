package internal

import "github.com/famcache/famcache-go/domain"

type Options struct {
	Host string
	Port int
}

func (o *Options) GetPort() int {
	return o.Port
}

func (o *Options) GetHost() string {
	return o.Host
}

func NewOptions(host string, port int) domain.Options {
	return &Options{
		Host: host,
		Port: port,
	}
}
