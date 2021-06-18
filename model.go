package main

// Nginx upstream representation
type NginxUpstream struct {
	Name string
	Ip   string
	Port int
}

// Ip filter configuration userd in locatin and server blocks
type NginxIpFilter struct {
	Allow bool
	Value string
}

// Nginx location definition
type NginxLocation struct {
	Path      string
	Return    int
	ProxyPass *NginxUpstream
	IpFilter  []NginxIpFilter
}

// Nginx block
type NginxServer struct {
	Listen     []string
	ServerName []string
	Root       string
	Location   map[string]NginxLocation
	IpFilter   []NginxIpFilter
}

// Nginx file model
type NginxModel struct {
	Upstreams map[string]NginxUpstream
	Servers   []NginxServer
}
