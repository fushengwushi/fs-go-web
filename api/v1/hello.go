package v1
import (
    "github.com/gin-gonic/gin"
    "fs-go-web/library"
    "fs-go-web/service"
)
func Hello(ctx *gin.Context)  {
    responseBody := library.NewResponseBody()
    defer library.RecoverResponse(ctx, responseBody)
    param := &service.HelloRequestParams{Ctx: ctx}
    ctx.BindJSON(param)
    service.Hello(param, responseBody)
}
