package profile

// Pluggable interface allow a plugin to be "plugged" in a profile.
// Is the common interface implemented by all plugins
type Pluggable interface {
	Name() string
	Renderable
}

// Renderable interface allow a plugin to render content in the shell loader
// and runner files
type Renderable interface {
	Render(profile Profile) string
}

// Setuppable interface allow a plugin to perform setup steps before rendering
type Setuppable interface {
	Setup(profile Profile) error
}