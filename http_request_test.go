package plus

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHTTPServerError(t *testing.T) {
	data := []struct {
		StatusCode int
		Result     bool
	}{
		{0, true},
		{http.StatusOK, false},
		{http.StatusCreated, false},
		{http.StatusNoContent, false},
		{http.StatusMovedPermanently, false},
		{http.StatusFound, false},
		{http.StatusBadRequest, false},
		{http.StatusForbidden, false},
		{http.StatusUnauthorized, false},
		{http.StatusUnsupportedMediaType, false},
		{http.StatusInternalServerError, true},
		{http.StatusBadGateway, true},
		{http.StatusGatewayTimeout, true},
		{http.StatusNotImplemented, true},
	}

	for d := range data {
		var resp http.Response
		if data[d].StatusCode != 0 {
			resp = http.Response{StatusCode: data[d].StatusCode}
		}
		assert.Equal(t, data[d].Result, IsHTTPServerError(&resp))
	}
}

func TestIsHTTPClientError(t *testing.T) {
	data := []struct {
		StatusCode int
		Result     bool
	}{
		{0, false},
		{http.StatusOK, false},
		{http.StatusCreated, false},
		{http.StatusNoContent, false},
		{http.StatusMovedPermanently, false},
		{http.StatusFound, false},
		{http.StatusBadRequest, true},
		{http.StatusForbidden, true},
		{http.StatusUnauthorized, true},
		{http.StatusUnsupportedMediaType, true},
		{http.StatusInternalServerError, false},
		{http.StatusBadGateway, false},
		{http.StatusGatewayTimeout, false},
		{http.StatusNotImplemented, false},
	}

	for d := range data {
		var resp http.Response
		if data[d].StatusCode != 0 {
			resp = http.Response{StatusCode: data[d].StatusCode}
		}
		assert.Equal(t, data[d].Result, IsHTTPClientError(&resp))
	}
}

func TestIsHTTPSuccessful(t *testing.T) {
	data := []struct {
		StatusCode int
		Result     bool
	}{
		{0, false},
		{http.StatusOK, true},
		{http.StatusCreated, true},
		{http.StatusNoContent, true},
		{http.StatusMovedPermanently, false},
		{http.StatusFound, false},
		{http.StatusBadRequest, false},
		{http.StatusForbidden, false},
		{http.StatusUnauthorized, false},
		{http.StatusUnsupportedMediaType, false},
		{http.StatusInternalServerError, false},
		{http.StatusBadGateway, false},
		{http.StatusGatewayTimeout, false},
		{http.StatusNotImplemented, false},
	}

	for d := range data {
		var resp http.Response
		if data[d].StatusCode != 0 {
			resp = http.Response{StatusCode: data[d].StatusCode}
		}
		assert.Equal(t, data[d].Result, IsHTTPSuccessful(&resp))
	}
}
