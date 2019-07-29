package profile

const (
	profilesDirectory   = "profiles"
	shellLoaderFilename = "load.sh"
	shellRunnerFilename = "run.sh"
)

// Profile holds information for a single profile
type Profile struct {
	// Name of the profile
	Name string
	// Location of the profile
	Location string
	// Plugins contains a map of Pluggable
	Plugins map[string]Pluggable
	// Shell is the shell to be used by profile
	Shell string

	// runLoaderPath is the path to be used for the profile run script
	runLoaderPath string
	// shellLoaderPath is the path to be used for the profile load script
	shellLoaderPath string
}

// Exists return where the profile exists
// Profile existance is determinated by profile Location existance
func (p Profile) Exists() bool {
	if p.Location == "" {
		return false
	}
	ok, err := exists(p.Location)
	if err != nil {
		return false
	}
	return ok
}
