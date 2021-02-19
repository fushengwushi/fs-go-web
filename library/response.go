package library
import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
)
type ResponseBody struct {
    Errno int32 `json:"errno"`
    Msg string `json:"msg"`
    Data interface{} `json:"data"`
}
func NewResponseBody() *ResponseBody {
    return &ResponseBody{
        Errno: ErrnoSuccess,
        Msg: GetErrMsg(ErrnoSuccess),
        Data: map[string]interface{}{},
    }
}
func (res *ResponseBody) SetData(data interface{})  {
    res.Data = data
}
func (res *ResponseBody) SetErrNo(errNo int32)  {
    res.Errno = errNo
}
func (res *ResponseBody) SetErrMsg(errMsg string)  {
    res.Msg = errMsg
}
func RecoverResponse(ctx *gin.Context, responseBody *ResponseBody)  {
    // panic
    if err := recover(); err != nil {
        responseBody.SetErrNo(ErrnoUnknown)
    }
    resp, err := json.Marshal(responseBody)
    if err != nil {
        ctx.Data(http.StatusOK, "application/json;charset=utf-8", []byte(`{"errno":1,"msg":"unknown"}`))
    } else {
        ctx.Data(http.StatusOK, "application/json;charset=utf-8", resp)
    }
    return
}
