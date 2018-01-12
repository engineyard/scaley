package group

type Server struct {
	ID              string `yaml:"id"`
	PrivateHostname string
}

type Group struct {
	Name             string    `yaml:"name"`
	PermanentServers []*Server `yaml:"permanent_servers"`
	ScalingServers   []*Server `yaml:"scaling_servers"`
	ScalingScript    string    `yaml:"scaling_script"`
}
