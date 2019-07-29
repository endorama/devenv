package profile

import (
	"os"
	"path"
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

func (p BinPlugin) Setup(profile Profile) error {
	return os.MkdirAll(path.Join(profile.Location, p.Name()), 0750)
}
