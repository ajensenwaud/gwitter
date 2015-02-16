package gwitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiError struct {
	StatusCode int
	Header     http.Header
	Body       string
	Decoded    TwitterErrorResponse
}

func (err ApiError) Error() string {
	return fmt.Sprintf("Twitter returned status %d: %s", err.StatusCode, err.Body)
}

type TwitterErrorResponse struct {
	Errors []TwitterError `json:"errors"`
}

// Satisfy Error interface
func (tr TwitterErrorResponse) Error() string {
	return tr.Errors[0].Message
}

type TwitterError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		return apiError(resp)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}

func apiError(resp *http.Response) *ApiError {
	str, _ := ioutil.ReadAll(resp.Body)

	var errResp TwitterErrorResponse
	_ = json.Unmarshal(str, &errResp)
	return &ApiError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Decoded:    errResp,
	}
}
