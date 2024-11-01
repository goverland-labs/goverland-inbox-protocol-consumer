package form

import (
	"net/http"

	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/response"
)

type Former interface {
	ParseAndValidate(request *http.Request) (Former, response.Error)
	ConvertToMap() map[string]interface{}
}
