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

	return r.StatusCode >= 500 && r.StatusCode <= 600
}
