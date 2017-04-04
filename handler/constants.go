package handler

var (
	STATUSES = map[string]status{
		"ok":                 {message: "200 OK", code: 200},
		"not_found":          {message: "404 NOT FOUND", code: 404},
		"bad_request":        {message: "400 BAD REQUEST", code: 400},
		"forbidden":          {message: "403 FORBIDDEN", code: 403},
		"method_not_allowed": {message: "405 METHOD NOT ALLOWED", code: 405},
		"not_supports":       {message: "505 HTTP VERSION NOT SUPPORTED", code: 505},
	}
	IMPLEMENTED_METHODS = []string{
		"GET",
		"HEAD",
	}
	CONTENT_TYPES = map[string]string{
		".css":  "text/css",
		".gif":  "image/gif",
		".html": "text/html",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".js":   "text/javascript",
		".json": "application/json",
		".txt":  "application/text",
		".png":  "image/png",
		".swf":  "application/x-shockwave-flash",
	}
)

const (
	STRING_SEPARATOR     string = "\r\n"
	HEADERS_END          string = STRING_SEPARATOR + STRING_SEPARATOR
	HTTP_VERSION         string = "HTTP/1.1"
	SERVER               string = "IliaBulavintsevServer"
	INDEX_FILE           string = "/index.html"
	ERROR_BODY_MIME_TYPE string = "text/html"
	DEFAULT_MIME_TYPE    string = "text/html"
)
