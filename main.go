package famcache

import (
	"github.com/famcache/famcache-go/domain"
	"github.com/famcache/famcache-go/internal"
)

func New(host string, port int) domain.Client {
	return internal.NewClient(host, port)
}
