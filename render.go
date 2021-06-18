package main

// template renderer
import (
	"bytes"
	"html/template"
)

func renderUpstreams(upstreams map[string]NginxUpstream) string {
	const upstreamTpl = `
{{ range . }}
upstream {{ .Name }} {
    server {{ .Ip }}:{{ .Port }};
}
{{  end }}
`
	var tpl bytes.Buffer
	t, err := template.New("renderUpstreams").Parse(upstreamTpl)
	check(err)
	err = t.Execute(&tpl, upstreams)
	check(err)
	return tpl.String()
}
