package config

type ServerConfiguration struct {
	TCP  Protocol
	UDP  Protocol
	HTTP Protocol
}

type Protocol struct {
	Active  bool
	Address string
}
