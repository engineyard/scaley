package eycore

import (
	"fmt"

	"github.com/ess/eygo"

	"github.com/engineyard/scaley/pkg/scaley"
)

type ServerService struct {
	upstream *eygo.ServerService
}

func NewServerService() *ServerService {
	return &ServerService{
		eygo.NewServerService(Driver),
	}
}

func (service *ServerService) Get(provisionedID string) (*scaley.Server, error) {
	return nil, fmt.Errorf("Unimplemented")
}

func (service *ServerService) Start(server *scaley.Server) error {
	return fmt.Errorf("Unimplemented")
}

func (service *ServerService) Stop(server *scaley.Server) error {
	return fmt.Errorf("Unimplemented")
}
