package main

// template renderer
import (
	"bytes"
	"html/template"
)

var upstreamTpl = env("UPSTREAM_TPL", `
{{ range . }}
upstream {{ .Name }} {
    server {{ .Ip }}:{{ .Port }};
}
{{  end }}
`)

func renderUpstreams(upstreams map[string]NginxUpstream) string {

	var tpl bytes.Buffer
	t, err := template.New("renderUpstreams").Parse(upstreamTpl)
	check(err)
	err = t.Execute(&tpl, upstreams)
	check(err)
	return tpl.String()
}

func renderServers(Servers []NginxServer) string {

	const nginxTpl = `
{{ range . }}
server {
  {{ range .Listen }} listen {{ . }};
  {{  end }}
  {{ range .ServerName }} server_name {{ . }};
  {{  end }}
  {{ if .Root }}
  root {{ .Root }};
  {{  end }}

  {{ range .Location }}
  location {{.Path}} {
    {{ if .Return }} return {{ .Return }}; {{  end }}
    {{ if .ProxyPass }} proxy_pass http://{{.ProxyPass.Name}};  {{  end }}
    {{ if .IpFilter }}
    {{ range .IpFilter }} {{ if .Allow }} allow {{ else }} deny {{ end }} {{ .Value }};
    {{ end }}
    {{ end }}
  }
  {{  end }}
  {{ range .IpFilter }}
    {{ if .Allow }} allow {{ else }} deny {{ end }} {{ .Value }};
  {{  end }}
}
{{  end }}
`

	var tpl bytes.Buffer
	t, err := template.New("servers").Parse(nginxTpl)
	check(err)

	err = t.Execute(&tpl, Servers)
	check(err)
	return tpl.String()
}
