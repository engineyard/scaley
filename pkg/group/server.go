package group

import (
	"github.com/engineyard/eycore/servers"
)

type Server struct {
	ID       string `yaml:"id"`
	Instance *servers.Model
}

func (s *Server) AmazonID() string {
	return s.ID
}

func (s *Server) EngineYardID() int {
	return s.Instance.ID
}

func (s *Server) Hostname() string {
	return s.Instance.PrivateHostname
}
