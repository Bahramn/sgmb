package config

type ServerConfiguration struct {
	TCP Protocol
	UDP Protocol
}

type Protocol struct {
	Active bool
	Address string
}
