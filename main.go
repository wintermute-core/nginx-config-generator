package main

import (
	"fmt"
	"log"
	"os"
)

const InputFileKey = "INPUT_FILE"
const InputFileValue = "input.yaml"

const ListenIp4Key = "LISTEN_IP4"
const ListenIp4Value = "0.0.0.0"

const ListenIp6Key = "LISTEN_IP6"
const ListenIp6Value = "[::]"

const UpstreamIpKey = "UPSTREAM_IP"
const UpstreamIpValue = "127.0.0.1"

const RootPath = "/"
const IpFilterKey = "ipfilter"

const DefaultListenPortKey = "DEFAULT_PORT"
const DefaultListenPortValue = "7000"

const DefaultPathKey = "DEFAULT_PATH"
const DefaultPathValue = "/var/www/"

const DefaultServerNameKey = "DEFAULT_SERVER_NAME"
const DefaultServerNameValue = "_"

func main() {
	log.Println("Nginx config generator")
	inputFile := env(InputFileKey, InputFileValue)
	log.Printf("Loading input file %s", inputFile)
	input := parse(inputFile)

	// build internal application model
	nginxModel := NginxModel{
		Upstreams: map[string]NginxUpstream{},
	}

	// default nginx server
	createDefaultServer(&nginxModel)
	// extract upstream definitions
	createUpstreams(&input, &nginxModel)

	extractApplications(&input, &nginxModel)

	outputFile := env("OUTPUT_FILE", "nginx.conf")
	log.Printf("Rendering config to %s", outputFile)

	// render output
	upstreams := renderUpstreams(nginxModel.Upstreams)
	servers := renderServers(nginxModel.Servers)

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s\n%s\n", upstreams, servers))
	check(err)
}

// Define "root" location in server configuration, if is not explicitly defined by application
func defineSeverRootLocation(server *NginxServer, n *NginxUpstream) {
	_, ok := server.Location[RootPath]
	if ok {
		return
	}
	server.Location[RootPath] = NginxLocation{
		Path:      RootPath,
		ProxyPass: n,
	}
}

// copy ip filter configuration from root location to server level
func defineServerIpFilter(nginxServer *NginxServer) {
	location, ok := nginxServer.Location[RootPath]
	if ok { // copy ipfilter from location to server
		nginxServer.IpFilter = location.IpFilter
		location.IpFilter = location.IpFilter[:0]
		nginxServer.Location[RootPath] = location
	}
}

// copy from app definition Fqdn into server structure
func extractFqdn(v App) []string {
	var fqdns []string
	for _, server := range v.Fqdn {
		fqdns = append(fqdns, server)
	}
	return fqdns
}

// map application paths into nginx server path
func extractPaths(v App, nginxUpstream NginxUpstream, i Input) map[string]NginxLocation {
	var nginxLocations = make(map[string]NginxLocation)
	for path, pathConfig := range v.PathBasedAccessRestriction {
		location := NginxLocation{
			Path:      path,
			ProxyPass: &nginxUpstream,
		}
		for configKey, configValue := range pathConfig {
			if configKey != IpFilterKey {
				log.Fatalf("Unknwon path config key %v", configKey)
			}
			ipfilter := i.IpFilter[configValue]
			for _, f := range ipfilter {
				filter := NginxIpFilter{
					Allow: true,
					Value: f,
				}
				location.IpFilter = append(location.IpFilter, filter)
			}
			if len(location.IpFilter) != 0 {
				filter := NginxIpFilter{
					Allow: false,
					Value: "all",
				}
				location.IpFilter = append(location.IpFilter, filter)
			}
		}
		nginxLocations[path] = location
	}
	return nginxLocations
}

// map applications to nginx server definitions
func extractApplications(input *Input, nginxModel *NginxModel) {
	for k, v := range input.App {
		log.Printf("Processing application %s\n", k)
		nginxUpstream := nginxModel.Upstreams[k]

		catchAllConfig, ok := input.CatchAll[v.CatchAll]
		if !ok {
			log.Fatalf("Unable to find catch all config %s", v.CatchAll)
		}
		catchPort := catchAllConfig["port"]
		nginxServer := NginxServer{
			Listen: []string{fmt.Sprintf("%v:%v", env(ListenIp6Key, ListenIp6Value), catchPort),
				fmt.Sprintf("%v:%v", env(ListenIp4Key, ListenIp4Value), catchPort)},
			ServerName: extractFqdn(v),
			Location:   extractPaths(v, nginxUpstream, *input),
			IpFilter:   []NginxIpFilter{},
		}

		defineSeverRootLocation(&nginxServer, &nginxUpstream)
		defineServerIpFilter(&nginxServer)

		nginxModel.Servers = append(nginxModel.Servers, nginxServer)
	}
}

// define upstreams based on input app
func createUpstreams(i *Input, nginxModel *NginxModel) {
	for key, v := range i.App {
		nginxUpstream := NginxUpstream{
			Name: key,
			Ip:   env(UpstreamIpKey, UpstreamIpValue),
			Port: v.RuntimePort,
		}
		nginxModel.Upstreams[key] = nginxUpstream
	}
}

// creat default server
func createDefaultServer(nginxModel *NginxModel) {
	defaultServer := NginxServer{
		Listen: []string{
			fmt.Sprintf("%v:%v default_server ipv6only=on",
				env(ListenIp6Key, ListenIp6Value), env(DefaultListenPortKey, DefaultListenPortValue)),

			fmt.Sprintf("%v:%v default_server",
				env(ListenIp4Key, ListenIp4Value), env(DefaultListenPortKey, DefaultListenPortValue)),
		},
		ServerName: []string{env(DefaultServerNameKey, DefaultServerNameValue)},
		Root:       env(DefaultPathKey, DefaultPathValue),
		Location: map[string]NginxLocation{
			RootPath: {
				Path:      RootPath,
				Return:    503,
				ProxyPass: nil,
				IpFilter:  nil,
			},
		},
	}
	nginxModel.Servers = append(nginxModel.Servers, defaultServer)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func env(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}
