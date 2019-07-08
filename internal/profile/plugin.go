package profile

type Pluggable interface {
	Name() string
	Render(profile Profile) string
}

// Configurable interface allow a plugin to load configuration from the profile
// folder
type Configurable interface {
	Config() interface{}
	ConfigFile(profileLocation string) string
	LoadConfig(profileLocation string) error
}
