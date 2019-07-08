package profile

import (
	"strings"
)

type AwsPlugin struct {
	Pluggable
}

func NewAwsPlugin() *AwsPlugin {
	return &AwsPlugin{}
}

func (p AwsPlugin) Name() string {
	return "aws"
}

func (p AwsPlugin) Render(profile Profile) string {
	sb := strings.Builder{}
	sb.WriteString("export AWS_CONFIG_FILE=" + profile.Location + "/aws/config\n")
	sb.WriteString("export AWS_SHARED_CREDENTIALS_FILE=" + profile.Location + "/aws/credentials\n")
	return sb.String()
}
