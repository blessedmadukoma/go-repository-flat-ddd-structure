package messages

import (
	"net/http"
)

// ResponseFormat handles the structure for returning response
type ResponseFormat struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Error   []string    `json:"error"`
	Message string      `json:"message"`
}

// Response create new instance of ResponseFormat
func Response(code int, res ResponseFormat) (int, ResponseFormat) {

	if code == 0 {
		code = 200
	}

	if !res.Status {
		res.Status = code < http.StatusBadRequest
	}

	if res.Data == nil {
		res.Data = make(map[string]interface{})
	}
	if res.Error == nil {
		res.Error = make([]string, 0)
	}
	if res.Message == "" {
		res.Message = OperationWasSuccessful
		if code == http.StatusNotFound {
			res.Message = NotFound
		} else if code >= http.StatusBadRequest {
			res.Message = SomethingWentWrong
		}
	}

	return code, res
}
