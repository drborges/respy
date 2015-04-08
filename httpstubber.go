package respy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	StatusOK                  = Replies{code: 200}
	StatusCreated             = Replies{code: 201}
	StatusBadRequest          = Replies{code: 400}
	StatusNotFound            = Replies{code: 404}
	StatusInternalServerError = Replies{code: 500}
)

type Replies struct {
	code    int
	body    string
	headers http.Header
}

func (r Replies) Header(key, value string) Replies {
	if r.headers == nil {
		r.headers = make(http.Header)
	}
	r.headers.Add(key, value)
	return r
}

func (r Replies) Body(json string) Replies {
	r.body = json
	return r
}

func (r Replies) Reply() (*httptest.Server, *http.Client) {
	return reply(r.code, r.body, r.headers)
}

// Code reused from http://keighl.com/post/mocking-http-responses-in-golang
func reply(code int, body string, headers http.Header) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range headers {
			w.Header().Add(key, value[0])
		}
		w.WriteHeader(code)
		fmt.Fprint(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client := &http.Client{Transport: transport}

	return server, client
}
