// Package scaley provides the domain objects, interfaces, and workflows for
// scaling server groups on Engine Yard Cloud.
package scaley

// Group is a representation of a scaling group.
type Group struct {
	Name           string   `yaml:"name"`
	ScalingServers []string `yaml:"scaling_servers"`
	ScalingScript  string   `yaml:"scaling_script"`
	StopScript     string   `yaml:"stop_script"`
	Strategy       string   `yaml:"strategy"`
}

// Server is a representation of a server in an Engine Yard Cloud environment.
type Server struct {
	ID            int
	ProvisionedID string
	State         string
	EnvironmentID string
}

// Environment is a representation of an Engine Yard Cloud environment.
type Environment struct {
	ID   string
	Name string
}
