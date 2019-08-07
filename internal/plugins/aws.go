package profile

import (
	"strings"
)

const AwsPluginName = "aws"

type AwsPlugin struct {}

func NewAwsPlugin() *AwsPlugin {
	return &AwsPlugin{}
}

func (p AwsPlugin) Name() string {
	return AwsPluginName
}

func (p AwsPlugin) Render(profileName, profileLocation string) string {
	sb := strings.Builder{}
	sb.WriteString("export AWS_CONFIG_FILE=" + profileLocation + "/aws/config\n")
	sb.WriteString("export AWS_SHARED_CREDENTIALS_FILE=" + profileLocation + "/aws/credentials\n")
	return sb.String()
}
