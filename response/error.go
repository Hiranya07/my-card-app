package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// errorMessage - Struct holds Error needs to send as a part of communication.
type errorMessage struct {
	cause      error
	code       string
	message    string
	statusCode int
}

// NewError - Construct a new error object to send as a part of communication.
func NewError(cause error, code string, message string, statusCode int) error {
	return errorMessage{cause: cause, code: code, message: message, statusCode: statusCode}
}

// Error - Message for the error.
func (e errorMessage) Error() string {
	return e.message
}

// response - error as a part of communication, With or without cause.
func (e errorMessage) response(w http.ResponseWriter, includeCause bool) error {
	message := map[string]string{"message": e.message, "code": e.code}

	if includeCause {
		message["cause"] = fmt.Sprintf("%v", e.cause)
	}

	data, err := json.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "error_cant_marshall")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.statusCode)

	_, err = w.Write(data)
	if err != nil {
		return errors.Wrap(err, "error_cant_process_data")
	}

	return nil
}
