# respy

respy provides a simple API for building http response stubs for your tests.

# Getting respy

```bash
$ go get github.com/drborges/respy
```

# Usage

Stubbing response `headers` and `body`:

```go
import "github.com/drborges/respy"

func TestStatusCreatedWithJSONBodyAndLocationHeader(t *testing.T) {
	expectedResourceLocation := "http://localhost/resource/1"
	expectedResponseBody := `{"message":"Resource created successfuly"}`

	server, client := respy.StatusCreated.
		Header("Location", expectedResourceLocation).
		Header("Content-type", "application/json").
		Body(expectedResponseBody).
		Reply()
	defer server.Close()

	resp, _ := client.Get(server.URL)
	body, _ := ioutil.ReadAll(reader)
	responseBody := string(body)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "http://localhost/resource/1", resp.Header.Get("Location"))
	assert.Equal(t, expectedResponseBody, responseBody)
}
```

Checking received request data on the `Server`:
 
```go
import "github.com/drborges/respy"

func TestServerStoresRequestInformation(t *testing.T) {
	server, client := respy.StatusOK.Reply()
	defer server.Close()

	json := `{"user": "drborges"}`
	client.Post(server.URL, "application/json", strings.NewReader(json))

	assert.Equal(t, "application/json", server.ReceivedRequest.Header.Get("Content-type"))
	assert.Equal(t, json, server.ReceivedRequest.Body)
}
```

# Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, please include a test function that reproduces the issue, that will help a lot to reduce back and forth :~

# License

The MIT License (MIT)

Copyright (c) 2015 Diego da Rocha Borges

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.