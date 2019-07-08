package template

import (
	"fmt"
	"text/template"
)

var (
	runnerTemplate = fmt.Sprintf(`
%s

eval $@
`, commonTemplate)
)

// GetRunnerTemplate return parsed runner template
func GetRunnerTemplate() (*template.Template, error) {
	return template.New("runner").
		Parse(runnerTemplate)
}
