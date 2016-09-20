package plus

import (
	"net/http"
	"reflect"
)

// IsHTTPServerError checks whether the provided response indicates a server error
func IsHTTPServerError(r *http.Response) bool {
	if reflect.DeepEqual(http.Response{}, *r) {
		return true
	}

	return r.StatusCode >= 500 && r.StatusCode < 600
}

// IsHTTPClientError checks whether the provided response indicates a client error
func IsHTTPClientError(r *http.Response) bool {
	if reflect.DeepEqual(http.Response{}, *r) {
		return false
	}

	return r.StatusCode >= 400 && r.StatusCode < 500
}

// IsHTTPSuccessful checks whether the provided response indicates a successful response
func IsHTTPSuccessful(r *http.Response) bool {
	if reflect.DeepEqual(http.Response{}, *r) {
		return false
	}

	return r.StatusCode >= 200 && r.StatusCode < 300
}
