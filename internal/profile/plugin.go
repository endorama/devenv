package profile

type Pluggable interface {
	Name() string
	Render(profile Profile) string
}

type Configurable interface {
	Config() interface{}
	ConfigFile(profileLocation string) string
	LoadConfig(profileLocation string) error
}
