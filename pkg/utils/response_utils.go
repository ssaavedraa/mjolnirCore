package utils

import (
	"hex/mjolnir-core/pkg/utils/logging"

	"github.com/gin-gonic/gin"
)

type ResponseFields map[string]interface{}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// deprecate in future
func ConvertToResponse(model interface{}, fields ResponseFields) interface{} {
	switch model := model.(type) {
	case gin.H:
		for key, value := range fields {
			model[key] = value
		}

		return model
	default:
		response := gin.H{}

		for key, value := range fields {
			response[key] = value
		}

		return response
	}
}

func RespondWithError(c *gin.Context, code int, message string, err error) {
	logging.Error(code, err)

	c.JSON(code, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
