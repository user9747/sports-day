package errorhandler

import (
	"encoding/json"
	"sports-day/internal/utils"

	"net/http"

	"golang.org/x/exp/slog"
)

const GenericFailureMessage = "something went wrong, please check back in an hour"

// Recovery handles the panic happening on any function, this is to be called by defer in functions
func Recovery(w http.ResponseWriter, request *http.Request, httpStatusCode int) {

	if r := recover(); r != nil {
		msg, ok := r.(string)
		if ok {
			// string message passed - no need to report to sentry, and use httpStatusCode
			CustomError(w, httpStatusCode, msg)
		} else {
			if !utils.InArr(httpStatusCode, []int{http.StatusUnauthorized, http.StatusForbidden}) {
				// set error code to use 500, if not an authorised error code
				httpStatusCode = http.StatusInternalServerError
			}
			err, ok := r.(error)
			msg = GenericFailureMessage
			if ok {
				// if error object found report to sentry
				slog.Error("recovered: ", r)
				slog.Error(StackTrace(Wrap(err)))
				CustomError(w, httpStatusCode, GenericFailureMessage)
			} else {
				// when string or error cannot be recovered (rare case)
				CustomError(w, httpStatusCode, GenericFailureMessage)
			}
		}
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		slog.Error(err.Error())
	}
}

// CustomError returns an error message without reporting to sentry
func CustomError(w http.ResponseWriter, httpStatusCode int, msg string) {
	errJSON := map[string]interface{}{
		"status": false,
		"data":   map[string]interface{}{},
		"error":  msg,
	}
	respondwithJSON(w, httpStatusCode, errJSON)
}
