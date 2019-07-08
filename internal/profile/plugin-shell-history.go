package profile

import (
	"fmt"
	"path"
	"strings"
)

type ShellHistoryPlugin struct {
	Pluggable
}

func NewShellHistoryPlugin() *ShellHistoryPlugin {
	return &ShellHistoryPlugin{}
}

func (p ShellHistoryPlugin) Name() string {
	return "shell-history"
}

func (p ShellHistoryPlugin) Render(profile Profile) string {
	sb := strings.Builder{}
	shellName := strings.ToLower(path.Base(profile.Shell))
	histFile := profile.Location + "/" + shellName + "-history"
	sb.WriteString(fmt.Sprintf("export HISTFILE=\"%s\"\n", histFile))
	return sb.String()
}
