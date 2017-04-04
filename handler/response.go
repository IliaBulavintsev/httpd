package handler

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type status struct {
	message string
	code    int
}

type response struct {
	status  *status
	headers map[string]string
}

func (r *response) set_status(status string) {
	if _, ok := STATUSES[status]; ok {
		*r.status = STATUSES[status]
	}
}

func (r *response) is_ok() bool {
	return *r.status == STATUSES["ok"]
}

func (handler Handler) write_response() {
	handler.write_string(HTTP_VERSION + " " + handler.response.status.message)
	handler.write_headers()
	handler.write_string("") // empty string after headers
	if handler.request.method != "HEAD" {
		handler.write_body()
	}
	fmt.Println(handler.request.method, " ", handler.get_path(), " ", handler.response.status.code)
}

func (handler *Handler) write_string(str string) {
	handler.Connection.Write([]byte(str + STRING_SEPARATOR))
}

func (handler *Handler) write_body() {
	if handler.response.is_ok() {
		handler.write_ok_body()
	} else {
		handler.write_error_body()
	}
}

func (handler *Handler) write_ok_body() {
	file, err := os.Open(handler.get_path())
	if err != nil {
		fmt.Println("Can't open file ", handler.get_path())
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_, read_err := reader.WriteTo(handler.Connection)
	if read_err != nil {
		fmt.Println("Some error on read or write file ", handler.get_path())
	}
}

func (handler *Handler) write_error_body() {
	body := []byte(handler.get_error_body())
	handler.Connection.Write(body)
}

func (handler *Handler) get_error_body() string {
	body := "<html><body><h1>"
	body += handler.response.status.message
	body += "</h1></body></html>"
	return body
}

func (handler Handler) write_headers() {
	handler.write_common_headers()
	handler.write_specific_headers()
}

func (handler Handler) write_common_headers() {
	handler.write_string("Date: " + time.Now().String())
	handler.write_string("Server: " + SERVER)
	handler.write_string("Connection: close")
}

func (handler Handler) write_specific_headers() {
	for key, value := range handler.response.headers {
		handler.write_string(key + ": " + value)
	}
}
