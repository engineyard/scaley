// Package scaley describes the high-level domain for scaling groups of
// servers on the Engine Yard Cloud platform.
package scaley

// Group is a representation of an autoscaling group
type Group struct {
	Name          string
	Servers       []Server
	ScalingScript string
	StopScript    string
	Strategy      string
	Environment   Environment
}

// Server is a representation of a server within a Group
type Server struct {
	ID            int
	ProvisionedID string
	State         int
	EnvironmentID string
}

// Environment is a representation of an environment on the Engine Yard Platform
type Environment struct {
	ID   string
	Name string
}

// Strategy provides the high-level functionality for scaling a group
type Strategy interface {
	Upscale() error
	Downscale() error
}

// GroupService provides retrieval functionality for groups
type GroupService interface {
	Get(string) (Group, error)
}

// ServerService provides retrieval functionality for servers
type ServerService interface {
	Get(string) (Server, error)
}

// EnvironmentService provides retrieval functionality for environments
type EnvironmentService interface {
	Get(string) (Environment, error)
}

// OpsService provides low-level functionality related to scaling groups.
//
// More than anything, an OpsService's focus is around starting/stopping
// servers, configuring environments, and so on.
type OpsService interface {
	Start(Server) error
	Stop(Server) error
	Configure(Environment) error
}
