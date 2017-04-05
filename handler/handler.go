package handler

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

type Handler struct {
	Connection net.Conn
	request    *request
	response   *response
	Factory    *Factory
}

func (handler *Handler) get_path() string {
	return handler.request.get_path()
}

func (handler *Handler) set_path(new_path string) {
	handler.request.set_path(new_path)
}

func (handler *Handler) set_header(key string, value string) {
	handler.response.headers[key] = value
}

func (handler *Handler) set_status(status string) {
	handler.response.set_status(status)
}

func (handler *Handler) read_request() {
	buffer := make([]byte, 1024)
	_, err := handler.Connection.Read(buffer)
	if err != nil {
		fmt.Println("Read request error ", err)
	}
	raw_request := string(buffer[:bytes.Index(buffer, []byte{0})])
	start_string := strings.Split(raw_request, STRING_SEPARATOR)[0]
	fmt.Println(start_string)
	handler.parse_start_string(start_string)
}

func (handler *Handler) parse_start_string(start_string string) {
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
	handler.request.url = parsed_url
}

func (handler *Handler) process_request() {
	if !handler.response.is_ok() {
		handler.set_content_headers(nil)
		return
	}
	if !contains(IMPLEMENTED_METHODS, handler.request.method) {
		handler.set_status("method_not_allowed")
	} else {
		handler.preprocess_path()
	}
}

//preproccess path and check file errors
func (handler *Handler) preprocess_path() {
	handler.set_path(handler.Factory.root + handler.get_path())
	file_info := handler.check_path(false)
	if file_info != nil && file_info.IsDir() {
		handler.set_path(handler.get_path() + INDEX_FILE)
		file_info = handler.check_path(true)
	}
	handler.set_content_headers(file_info)
}

func (handler *Handler) check_path(is_dir bool) os.FileInfo {
	request_path := handler.get_path()
	clear_path := path.Clean(request_path)
	handler.set_path(clear_path)
	info, err := os.Stat(request_path)
	if err != nil {

		if os.IsNotExist(err) && !is_dir {
			handler.set_status("not_found")
		} else {
			handler.set_status("forbidden")
		}
	} else if !strings.Contains(clear_path, handler.Factory.root) {
		handler.set_status("forbidden")
	}
	return info
}

func (handler *Handler) set_content_headers(info os.FileInfo) {
	if handler.response.is_ok() {
		handler.set_header("Content-Length", strconv.Itoa(int(info.Size())))
		handler.set_header("Content-Type", handler.get_content_type())
	} else {
		handler.set_header("Content-Length", strconv.Itoa(len(handler.get_error_body())))
		handler.set_header("Content-Type", ERROR_BODY_MIME_TYPE)
	}
}

func contains(arr []string, value string) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}
	return false
}

func (handler *Handler) get_content_type() string {
	extension := ""
	request_path := handler.get_path()
	last_dot := strings.LastIndex(request_path, ".")
	if last_dot >= 0 {
		extension = request_path[last_dot:]
	}
	val, ok := CONTENT_TYPES[extension]
	if ok {
		return val
	} else {
		return DEFAULT_MIME_TYPE
	}
}

func (handler *Handler) clear() {
	handler.Factory = nil
	handler.Connection.Close()
}

func (handler *Handler) Handle() {
	handler.read_request()
	handler.process_request()
	handler.write_response()
	handler.clear()
}
