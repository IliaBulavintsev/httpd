package handler

import (
	"net/url"
)

type request struct {
	method string
	url    *url.URL
}

func (request *request) get_path() string {
	if request.url != nil {
		return request.url.Path
	} else {
		return ""
	}
}

func (request *request) set_path(new_path string) {
	request.url.Path = new_path
}
