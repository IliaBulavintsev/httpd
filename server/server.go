package server

import (
	"fmt"
	"net"

	"../handler"
)

type Server struct {
	port string
	host string
	handler.Storage
}

func (s *Server) Create(host string, port string, root string) {
	s.port = port
	s.host = host
	s.Storage.CreateStorage(root)

	fmt.Println("storage:::", s.Storage.Get_root())
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
		// handler := handler.Handler{Connection: connection, Storage: s.Storage}

		go handler.Handle()
	}
}
