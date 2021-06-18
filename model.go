package main

// Input model

type IpFilter []string

type CatchAll map[string]interface{}

type PathBasedMapping map[string]string

type App struct {
	CatchAll                   string                      `yaml:"catchall"`
	Fqdn                       []string                    `yaml:"fqdn"`
	RuntimePort                int                         `yaml:"runtime_port"`
	PathBasedAccessRestriction map[string]PathBasedMapping `yaml:"path_based_access_restriction"`
}

type Input struct {
	IpFilter map[string]IpFilter `yaml:ipfilter`
	CatchAll map[string]CatchAll `yaml:"catchall"`
	App      map[string]App      `yaml:"app"`
}

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
