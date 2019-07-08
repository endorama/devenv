package profile

import (
	"strings"
)

type BinPlugin struct {
	Pluggable
}

func NewBinPlugin() *BinPlugin {
	return &BinPlugin{}
}

func (p BinPlugin) Name() string {
	return "bin"
}

func (p BinPlugin) Render(profile Profile) string {
	sb := strings.Builder{}
	sb.WriteString("export PATH=" + profile.Location + "/bin:$PATH\n")
	return sb.String()
}
