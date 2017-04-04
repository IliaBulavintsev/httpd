package handler

import (
	"fmt"
	"net"
	"path/filepath"
)

type Storage struct {
	root string
}

func (s *Storage) CreateStorage(root string) {
	abs_root, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Storage create failed")
	}
	s.root = abs_root
	fmt.Println("root storage: ", s.root)
}

func (s *Storage) Get_root() string {
	return s.root
}

func (storage *Storage) CreateHandler(connection net.Conn) Handler {
	handler := Handler{}
	handler.request = new(request)
	handler.response = new(response)
	handler.response.status = new(status)
	handler.response.set_status("ok")
	handler.response.headers = map[string]string{}
	handler.Storage = storage
	handler.Connection = connection
	return handler
}
