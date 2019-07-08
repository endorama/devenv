package profile

type Pluggable interface {
	Name() string
	Render(profile Profile) string
}
