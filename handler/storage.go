package handler

import (
	"fmt"
	"net"
	"path/filepath"
)

type Factory struct {
	root string
}

func (f *Factory) CreateFactory(root string) {
	abs_root, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Storage create failed")
	}
	f.root = abs_root
}

func (f *Factory) Get_root() string {
	return f.root
}

func (f *Factory) CreateHandler(connection net.Conn) Handler {
	handler := Handler{}
	handler.request = new(request)
	handler.response = new(response)
	handler.response.status = new(status)
	handler.response.set_status("ok")
	handler.response.headers = map[string]string{}
	handler.Factory = f
	handler.Connection = connection
	return handler
}
