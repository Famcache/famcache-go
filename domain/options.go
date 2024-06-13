package domain

type Options interface {
	GetPort() int
	GetHost() string
}
