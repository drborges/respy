package respy

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

func bodyAsJSON(reader io.Reader) string {
	body, _ := ioutil.ReadAll(reader)
	return string(body)
}

func TestStatusBadRequest(t *testing.T) {
	server, client := StatusBadRequest.Reply()
	defer server.Close()

	resp, _ := client.Get(server.URL)

	assert.Equal(t, 400, resp.StatusCode)
}

func TestStatusOkWithBody(t *testing.T) {
	serverResponseBody := `{"user":"drborges"}`
	server, client := StatusOK.Body(serverResponseBody).Reply()
	defer server.Close()

	resp, _ := client.Get(server.URL)
	responseBody := bodyAsJSON(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, serverResponseBody, responseBody)
}

func TestStatusCreatedWithHeaders(t *testing.T) {
	resourceLocation := "http://localhost/resource/1"
	server, client := StatusCreated.Header("Location", resourceLocation).Reply()
	defer server.Close()

	resp, _ := client.Get(server.URL)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, resourceLocation, resp.Header.Get("Location"))
}

func TestStatusCreatedWithJSONBodyAndLocationHeader(t *testing.T) {
	expectedResourceLocation := "http://localhost/resource/1"
	expectedResponseBody := `{"message":"Resource created successfuly"}`

	server, client := StatusCreated.
		Header("Location", expectedResourceLocation).
		Header("Content-type", "application/json").
		Body(expectedResponseBody).
		Reply()
	defer server.Close()

	resp, _ := client.Get(server.URL)
	responseBody := bodyAsJSON(resp.Body)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "http://localhost/resource/1", resp.Header.Get("Location"))
	assert.Equal(t, expectedResponseBody, responseBody)
}
