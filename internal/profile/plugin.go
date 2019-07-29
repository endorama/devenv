package profile

// Pluggable interface allow a plugin to be "plugged" in a profile.
// Is the common interface implemented by all plugins
type Pluggable interface {
	Name() string
}