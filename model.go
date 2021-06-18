package main

// Internal model structures
type NginxUpstream struct {
	Name string
	Ip   string
	Port int
}

// Nginx file model
type NginxModel struct {
	Upstreams map[string]NginxUpstream
}
