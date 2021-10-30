package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_APIServer_HandleHello_ShouldReturnHelloWorld(t *testing.T) {
	server := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("Get", "/", nil)
	server.HandleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello World!")

}
