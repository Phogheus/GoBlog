package goblog_services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

type ErrorResponse struct {
	Error string
}

type HealthResponse struct {
	Code   int
	Status string
}

var (
	BadPathError   ErrorResponse = ErrorResponse{"Bad path."}
	BadMethodError ErrorResponse = ErrorResponse{"Bad method."}
)

func HandleRequests(writer http.ResponseWriter, request *http.Request) {
	var head string
	head, request.URL.Path = ShiftPath(request.URL.Path)

	switch head {
	case "":
		serveHomePage(writer)
	case "health":
		if request.Method == "GET" {
			serveHealthStatus(writer)
		} else {
			json.NewEncoder(writer).Encode(BadMethodError)
		}
	case "blog":
		HandleBlogRequest(writer, request)
	case "favicon.ico":
		// TODO
	default:
		json.NewEncoder(writer).Encode(BadPathError)
	}
}

// Discovered at https://benhoyt.com/writings/go-routing/#shiftpath
// Original source at https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)            // Prepend "/" to 'p' and clean - Clean will ensure proper formatting, allowing a "//" to become "/"
	i := strings.Index(p[1:], "/") + 1 // Find the next index of "/" after the prepended one (the trail)

	// If index is <= 0, there is no further trail
	if i <= 0 {
		return p[1:], "/"
	}

	// Slice path into head and trail (the rest)
	return p[1:i], p[i:]
}

func serveHomePage(writer http.ResponseWriter) {
	// Placeholder for some sort of frontend
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}

func serveHealthStatus(writer http.ResponseWriter) {
	// What kind of statuses would be of benefit here?
	json.NewEncoder(writer).Encode(HealthResponse{0, "Healthy"})
}
