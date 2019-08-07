package profile

import (
	"os"
	"path"
	"strings"
)

const BinPluginName = "bin"

type BinPlugin struct {}

func NewBinPlugin() *BinPlugin {
	return &BinPlugin{}
}

func (p BinPlugin) Name() string {
	return BinPluginName
}

func (p BinPlugin) Render(profileName, profileLocation string) string {
	sb := strings.Builder{}
	sb.WriteString("export PATH=" + profileLocation + "/bin:$PATH\n")
	return sb.String()
}

func (p BinPlugin) Setup(profileLocation string) error {
	return os.MkdirAll(path.Join(profileLocation, BinPluginName), 0750)
}
