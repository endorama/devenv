package profile

import (
	"fmt"
	"strings"
)

const ShellHistoryPluginName = "shell-history"

type ShellHistoryPlugin struct {}

func NewShellHistoryPlugin() *ShellHistoryPlugin {
	return &ShellHistoryPlugin{}
}

func (p ShellHistoryPlugin) Name() string {
	return ShellHistoryPluginName
}

func (p ShellHistoryPlugin) Render(profileName, profileLocation string) string {
	sb := strings.Builder{}
	histFile := profileLocation + "/" + ShellHistoryPluginName
	sb.WriteString(fmt.Sprintf("export HISTFILE=\"%s\"\n", histFile))
	return sb.String()
}
