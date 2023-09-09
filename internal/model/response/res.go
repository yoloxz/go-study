package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(res Response, c *gin.Context) {
	c.JSON(res.Code, res)
}

func OK(result interface{}, c *gin.Context) Response {
	return Response{
		Code: 200,
		Msg:  "success",
		Data: result,
	}
}

func Fail(code int, msg string, c *gin.Context) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
