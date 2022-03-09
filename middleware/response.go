package middleware

import "github.com/gin-gonic/gin"

func Response(response int, error string, appid string, svcid string, controller string, action string, data interface{}, c *gin.Context) {

	switch action {
	case "POST":
		action = "add"
	case "GET":
		action = "view"
	case "PUT":
		action = "edit"
	case "DELETE":
		action = "delete"
	default:
		action = action
	}

	res := struct {
		Response   int         `json:"response"`
		Error      string      `json:"error"`
		Appid      string      `json:"appid"`
		Svcid      string      `json:"svcid"`
		Controller string      `json:"controller"`
		Action     string      `json:"action"`
		Result     interface{} `json:"result"`
	}{
		Response:   response,
		Error:      error,
		Appid:      appid,
		Svcid:      svcid,
		Controller: controller,
		Action:     action,
		Result:     data,
	}
	c.JSON(response, res)

}
