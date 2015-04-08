package respy

import (
	"net/http"
	"net/http/httptest"
)

type requestInfo struct {
	Body string
	*http.Request
}

type Server struct {
	*httptest.Server
	ReceivedRequest requestInfo
}
