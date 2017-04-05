package server

import (
	"fmt"
	"net"

	"../handler"
)

type Server struct {
	port string
	host string
	handler.Factory
}

func (s *Server) Create(host string, port string, root string) {
	s.port = port
	s.host = host
	s.Factory.CreateFactory(root)
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		fmt.Println("Start server failed ", err)
		return
	} else {
		fmt.Println("Server start ", listener.Addr())
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handler := s.CreateHandler(connection)
		fmt.Println("CREATE Handler!!\n\n")

		go handler.Handle()
	}
}
