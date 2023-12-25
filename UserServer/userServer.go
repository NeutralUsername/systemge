package UserServer

import (
	"ServerUtilities"
)

type UserServer struct {
	Base *ServerUtilities.ServerBase
}

func CreateUserServer() *UserServer {
	server := &UserServer{
		Base: ServerUtilities.CreateServerBase(),
	}
	return server
}
