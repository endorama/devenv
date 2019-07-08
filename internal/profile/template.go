package profile

import (
	"fmt"
	"strings"
	"text/template"
)

var (
	commonTemplate = `
#!/bin/bash
#
# This file has been automatically generated with devenv
# Please remember that running 'devenv rehash' will overwrite this file :)

export DEVENV_ACTIVE_PROFILE='{{.Name}}'
export DEVENV_ACTIVE_PROFILE_PATH='{{.Location}}'

# plugin BEGIN ##################
{{range $key, $value := .Plugins}}{{if $value}}
# plugin: {{$key}}
{{.Render $}}{{end}}{{end}}

# plugin END ####################`

	shellLoaderTemplate = fmt.Sprintf(`
%s

exec {{.Shell}} -l
`, commonTemplate)

	runTemplate = fmt.Sprintf(`
%s

eval $@
`, commonTemplate)
)

type RunnerTemplate template.Template

func templateRenderPlugin(profile Profile, pluginName string) string {
	var sb strings.Builder
	switch pluginName {
	default:
		sb.WriteString("__devenv_plugin__" + pluginName + "__generate_loader")
	}
	return sb.String()
}

func getShellLoaderTemplate() (*template.Template, error) {
	return template.New("shell-loader").
		Funcs(template.FuncMap{
			"renderPlugin": templateRenderPlugin,
		}).
		Parse(shellLoaderTemplate)
}

func getRunnerTemplate() (*template.Template, error) {
	return template.New("runner").
		Funcs(template.FuncMap{
			"renderPlugin": templateRenderPlugin,
		}).
		Parse(runTemplate)
}
