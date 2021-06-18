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
	}
	output := renderUpstreams(upstreams)
	if !strings.Contains(output, "Potato123") {
		t.Error("Upstream name not rendered")
	}
	if !strings.Contains(output, "6.6.6.6") {
		t.Error("Upstream IP not rendered")
	}
	if !strings.Contains(output, "666") {
		t.Error("Upstream Port not rendered")
	}
}
