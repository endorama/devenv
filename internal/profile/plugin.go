package profile

type Pluggable interface {
	Name() string
	Render(profile Profile) string
}

// Generator interface allow a plugin to generate content before rendering
type Generator interface {
	Generate(profile Profile) error
}

// Configurable interface allow a plugin to load configuration from the profile
// folder
type Configurable interface {
	Config() interface{}
	ConfigFile(profileLocation string) string
	LoadConfig(profileLocation string) error
}
