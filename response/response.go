package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

func Response(w http.ResponseWriter, statusCode int, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return errors.Wrap(err, "error_cant_marshall")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	_, err = w.Write(data)
	if err != nil {
		return errors.Wrap(err, "error_cant_process_data")
	}

	return nil
}

// Error - Send error back to caller.
func Error(ctx context.Context, w http.ResponseWriter, err error, includeCause bool) {
	e, ok := err.(errorMessage)
	if !ok {
		e = errorMessage{cause: err, code: "error_internal_server_error", message: err.Error(), statusCode: http.StatusInternalServerError}
	}

	err = e.response(w, includeCause)
	if err != nil {
		log.Error(ctx, "response-error", "Unable to send error message %v", err)
	}
}
