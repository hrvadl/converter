package handlers

import "encoding/json"

// NewErrResponse constructs the JSON encoded response,
// which represents request failure from API.
// NOTE: it expects error to be not nil value.
func NewErrResponse(err error) []byte {
	bytes, _ := json.Marshal(ErrorResponse{
		Err:     err.Error(),
		Success: false,
	})
	return bytes
}

// ErrorResponse struct is a JSON encoded response,
// which represents request failure from API.
// NOTE: it expects error to be not nil value.
type ErrorResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"error"`
}

// NewEmptyResponse constructs the JSON encoded response,
// which represents request success from API, but doesn't have any
// data, so it omits data field in reponse.
func NewEmptyResponse(msg string) []byte {
	bytes, _ := json.Marshal(EmptyResponse{
		Msg:     msg,
		Success: true,
	})
	return bytes
}

// EmptyResponse struct is a JSON encoded response,
// which represents request success from API, but doesn't have any
// data, so it omits data field in reponse.
type EmptyResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"message,omitempty"`
}

// NewEmptyResponse constructs the JSON encoded response,
// which represents request success from API.
// NOTE: if data is nil then it will be omitter in the result.
func NewSuccessResponse[T any](msg string, data T) []byte {
	bytes, _ := json.Marshal(Response[T]{
		EmptyResponse: EmptyResponse{
			Msg:     msg,
			Success: true,
		},
		Data: data,
	})
	return bytes
}

// Response struct is a JSON encoded response,
// which represents request success from API.
// NOTE: if data is nil then it will be omitter in the result.
type Response[T any] struct {
	EmptyResponse
	Data T `json:"data,omitempty"`
}
