package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerBase(t *testing.T) {

	// serv := NewServer(config.NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	HandleBase(rec, req)
	assert.Equal(t, rec.Body.String(), "BASE RESPONSE")
}
