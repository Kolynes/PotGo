package response

import "net/http"

type Response struct {
	Status  int
	Body    []byte
	Headers http.Header
}
