package profile

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	utils "github.com/endorama/devenv/internal/utils"
)

const (
	TmuxPluginName = "tmux"
)

type TmuxPlugin struct {}

func NewTmuxPlugin() *TmuxPlugin {
	return &TmuxPlugin{}
}

func (p TmuxPlugin) Name() string {
	return TmuxPluginName
}

func (p TmuxPlugin) Render(profileName, profileLocation string) string {
	return fmt.Sprintf("export TMUX_SOCKET_NAME=\"%s.%s\"\n", TmuxPluginName, profileName)
}

func (p *TmuxPlugin) Generate(profileLocation string) error {
	tmux := strings.Builder{}
	tmux.WriteString("#!/usr/bin/env bash\n")
	tmux.WriteString("exec ")

	// we need to use the absolute path or will end up in a loop
	systemTmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		return fmt.Errorf("cannot lookup tmux path: %w", err)
	}
	tmux.WriteString(systemTmuxPath)

	tmux.WriteString(" -S \"/tmp/devenv_$TMUX_SOCKET_NAME\" ")
	tmux.WriteString("$@")

	binFilePath := path.Join(profileLocation, "bin", "tmux")
	utils.PersistFile(binFilePath, tmux.String())
	os.Chmod(binFilePath, 0700)

	return nil
}
