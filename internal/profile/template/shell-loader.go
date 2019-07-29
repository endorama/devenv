package template

import (
	"fmt"
	"text/template"
)

var (
	shellLoaderTemplate = fmt.Sprintf(`
%s

exec {{.Shell}} -l
`, commonTemplate)
)

// GetShellLoaderTemplate return parsed shell loader template
func GetShellLoaderTemplate() (*template.Template, error) {
	return template.New("shell-loader").
		Parse(shellLoaderTemplate)
}
