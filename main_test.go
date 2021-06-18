package main

import "testing"

func TestDefaultNginx(t *testing.T) {

	nginxModel := NginxModel{}

	createDefaultServer(&nginxModel)
	if len(nginxModel.Servers) == 0 {
		t.Error("Default server not created")
	}
	server := nginxModel.Servers[0]
	if server.Root != "/var/www/" {
		t.Error("Invalid root config")
	}

	if len(server.Location) != 1 {
		t.Error("Invalid location configuration")
	}
}

func TestUpstreamExtraction(t *testing.T) {
	input := Input{
		App: map[string]App{
			"Tomato": App{},
		},
	}
	nginxModel := NginxModel{
		Upstreams: map[string]NginxUpstream{},
	}

	createUpstreams(&input, &nginxModel)

	if len(nginxModel.Upstreams) != 1 {
		t.Error("Invalid upstream config")
	}

	upstream, ok := nginxModel.Upstreams["Tomato"]
	if !ok {
		t.Error("Expected upstream not created")
	}

	if upstream.Name != "Tomato" {
		t.Error("Invalid upstream name")
	}
}

func TestAppExtraction(t *testing.T) {
	input := Input{
		IpFilter: map[string]IpFilter{
			"Tomato": {
				"6.6.6.6",
			},
		},
		CatchAll: map[string]CatchAll{
			"Potato1": {
				"port": 5555,
			},
			"Potato2": {
				"port": 12345,
			},
		},
		App: map[string]App{
			"Tomato": App{
				CatchAll:    "Potato1",
				Fqdn:        []string{"service1", "service2"},
				RuntimePort: 666,
				PathBasedAccessRestriction: map[string]PathBasedMapping{
					"/cucumber": {
						IpFilterKey: "Tomato",
					},
				},
			},
		},
	}
	nginxModel := NginxModel{
		Upstreams: map[string]NginxUpstream{},
	}

	createUpstreams(&input, &nginxModel)
	extractApplications(&input, &nginxModel)

	if len(nginxModel.Servers) != 1 {
		t.Fatalf("Failed to map servers, got %d", len(nginxModel.Servers))
	}
	server := nginxModel.Servers[0]
	if len(server.ServerName) != 2 {
		t.Fatalf("Failed to map server name, got %d", len(server.ServerName))
	}

	cucumber, ok := server.Location["/cucumber"]
	if !ok {
		t.Fatal("Explicit location /cucumber not mapped")
	}

	if cucumber.ProxyPass.Name != "Tomato" {
		t.Fatalf("Proxy pass not mapped, got %s", cucumber.ProxyPass.Name)
	}

	root, ok := server.Location["/"]
	if !ok {
		t.Fatal("Default location / not mapped")
	}

	if root.ProxyPass.Name != "Tomato" {
		t.Fatalf("Proxy pass not mapped, got %s", root.ProxyPass.Name)
	}

}
