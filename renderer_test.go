package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestUpstreamRender(t *testing.T) {
	var upstreams = map[string]NginxUpstream{
		"Potato123": {
			Name: "Potato123",
			Ip:   "6.6.6.6",
			Port: 666,
		},
		"Tomato123": {
			Name: "Tomato123",
			Ip:   "1.2.3.4",
			Port: 123,
		},
	}
	output := renderUpstreams(upstreams)
	if !strings.Contains(output, "Potato123") {
		t.Error("Upstream name not rendered")
	}
	if !strings.Contains(output, "Tomato123") {
		t.Error("Upstream name not rendered")
	}
	if !strings.Contains(output, "6.6.6.6") {
		t.Error("Upstream IP not rendered")
	}
	if !strings.Contains(output, "1.2.3.4") {
		t.Error("Upstream IP not rendered")
	}
	if !strings.Contains(output, "666") {
		t.Error("Upstream Port not rendered")
	}
	if !strings.Contains(output, "123") {
		t.Error("Upstream Port not rendered")
	}
}

func TestServersRenderer(t *testing.T) {
	var nginxServers = []NginxServer{
		{
			Listen:     []string{"1.2.3.4:666"},
			ServerName: []string{"potato.com", "tomato.net"},
			Root:       "/potato/tomato",
			Location: map[string]NginxLocation{
				"/": {
					Path:   "/",
					Return: 6969,
					ProxyPass: &NginxUpstream{
						Name: "upstream666",
						Ip:   "9.8.7.6",
						Port: 5432,
					},
				},
			},
			IpFilter: []NginxIpFilter{
				{
					Allow: true,
					Value: "5.6.7.8",
				},
			},
		},
	}

	output := renderServers(nginxServers)
	fmt.Printf("%s\n", output)
}
