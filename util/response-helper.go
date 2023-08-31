package util

import "net/http"

type response struct {
	status int
	writer http.ResponseWriter
}

func (r *response) Writer(w http.ResponseWriter) *response {
	r.writer = w
	return r
}

func (r *response) Status422() *response {
	r.status = http.StatusUnprocessableEntity

	return r
}

func (r *response) Status403() *response {
	r.status = http.StatusForbidden

	return r
}

func (r *response) Status500() *response {
	r.status = http.StatusInternalServerError
	return r
}

func (r *response) Status(status ...int) *response {
	statusCode := http.StatusAccepted

	if len(status) > 0 {
		statusCode = status[0]
	}

	r.status = statusCode

	return r
}

func (r *response) Data(data any) {
	r.writer.WriteHeader(r.status)
	JsonEncoder(r.writer, data)
}

var responseStruct = response{
	status: http.StatusAccepted,
}

var Res = &responseStruct
