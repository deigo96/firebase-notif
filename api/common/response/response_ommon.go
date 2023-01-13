package response

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildSuccsessResponse(message string, status bool, data interface{}) Response {
	res := Response{
		Message: message,
		Status:  status,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, data interface{}) Response {
	res := Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
	return res
}
