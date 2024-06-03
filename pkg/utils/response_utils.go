package utils

import "github.com/gin-gonic/gin"

type ResponseFields map[string]interface{}

func ConvertToResponse (model interface{}, fields ResponseFields) interface{} {
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