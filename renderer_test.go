package main

import (
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
		t.Fatal("Upstream name not rendered")
	}
	if !strings.Contains(output, "Tomato123") {
		t.Fatal("Upstream name not rendered")
	}
	if !strings.Contains(output, "6.6.6.6") {
		t.Fatal("Upstream IP not rendered")
	}
	if !strings.Contains(output, "1.2.3.4") {
		t.Fatal("Upstream IP not rendered")
	}
	if !strings.Contains(output, "666") {
		t.Fatal("Upstream Port not rendered")
	}
	if !strings.Contains(output, "123") {
		t.Fatal("Upstream Port not rendered")
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
					IpFilter: []NginxIpFilter{
						{
							Allow: true,
							Value: "4.4.4.4",
						},
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

	if !strings.Contains(output, "1.2.3.4:666") {
		t.Fatal("Server listen line not rendered")
	}

	if !strings.Contains(output, "potato.com") {
		t.Fatal("Server name not rendered")
	}

	if !strings.Contains(output, "/potato/tomato") {
		t.Fatal("Root path not rendered")
	}

	if !strings.Contains(output, "upstream666") {
		t.Fatal("Upstream not included")
	}

	if !strings.Contains(output, "4.4.4.4") {
		t.Fatal("Ip filter for location not included")
	}

	if !strings.Contains(output, "5.6.7.8") {
		t.Fatal("Ip filter for server not included")
	}
}
