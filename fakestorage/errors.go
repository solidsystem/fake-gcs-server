// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fakestorage

import (
	"net/http"
)

// ServerError is for signalling the same error messages that the real server returns,
// with http statuscode and text.
type ServerError struct {
	StatusCode int
	Msg        string
}

func (se *ServerError) Error() string {
	return se.Msg
}

var PreconditionFailed = &ServerError{StatusCode: http.StatusPreconditionFailed, Msg: "Precondition Failed"}
