package handler

import (
	"net"
)

type Handler struct {
	connection net.Conn
	request
	response
	storage Storage
}

func (handler *Handler) Handle() {
	handler.read_request()
	handler.write_response()
	handler.clear()
}
