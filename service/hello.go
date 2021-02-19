package service

import (
    "fs-go-web/library"
    "fs-go-web/library/log"
    "github.com/gin-gonic/gin"
)
// post请求参数结构体
type HelloRequestParams struct {
    Text string `json:"text"`
    Ctx *gin.Context
}
func Hello(param *HelloRequestParams, responseBody *library.ResponseBody)  {
    // 获取get参数
    text := param.Ctx.DefaultQuery("text", "")
    if len(text) <= 0 {
        text = param.Text
        log.Info(map[string]interface{}{"from post": text})
    } else {
        log.Info(map[string]interface{}{"from get": text})
    }
    responseBody.SetData(text)
    return
}
