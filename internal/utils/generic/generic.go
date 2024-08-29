package gen

import "net/http"

type GenericResponse struct {
	Data   interface{}
	Status int
}

func GetResponse(data interface{}, status int) *GenericResponse {
	return &GenericResponse{
		Data:   data,
		Status: status,
	}
}
func HandleError(err error, message interface{}) *GenericResponse {
	return GetResponse(GetErrorJson(message, http.StatusBadRequest), http.StatusOK)
}
func GetErrorJson(message interface{}, statuscode int) map[string]interface{} {
	return map[string]interface{}{"message": message, "success": false, "code": statuscode}
}
