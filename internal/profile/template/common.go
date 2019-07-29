package template

var (
	commonTemplate = `#!/bin/bash
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
)
