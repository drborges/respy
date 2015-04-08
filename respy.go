package respy

import (
	"fmt"
	"net/url"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
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

func (r Replies) Reply() (*Server, *http.Client) {
	return reply(r.code, r.body, r.headers)
}

// Code reused from http://keighl.com/post/mocking-http-responses-in-golang
func reply(code int, body string, headers http.Header) (*Server, *http.Client) {
	wrapper := &Server{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range headers {
			w.Header().Add(key, value[0])
		}

		reqBody, _ := ioutil.ReadAll(r.Body)
		wrapper.ReceivedRequest = requestInfo{
			Request: r,
			Body: string(reqBody),
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

	wrapper.Server = server
	return wrapper, client
}
