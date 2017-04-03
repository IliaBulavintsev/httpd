package handler

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

func (handler *http_handler) read_request() {
	buffer := make([]byte, 1024)
	_, err := handler.connection.Read(buffer)
	if err != nil {
		fmt.Println("Read request error ", err)
	}
	raw_request := string(buffer[:bytes.Index(buffer, []byte{0})])
	start_string := strings.Split(raw_request, STRING_SEPARATOR)[0]
	handler.parse_start_string(start_string)
}

func (handler *http_handler) parse_start_string(start_string string) {
	splited_string := strings.Split(start_string, " ")
	if len(splited_string) != 3 {
		handler.set_status("bad_request")
		return
	}
	handler.request.method = splited_string[0]
	parsed_url, err := url.Parse(splited_string[1])
	if err != nil || !strings.HasPrefix(splited_string[2], "HTTP/") {
		handler.set_status("bad_request")
	}
	handler.request.request_url = parsed_url
}
