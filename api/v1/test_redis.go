package v1

import (
	"fs-go-web/library"
	"fs-go-web/service"
	"github.com/gin-gonic/gin"
)

func TestRedis(ctx *gin.Context)  {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(ctx, responseBody)

	service.TestRedis(ctx, responseBody)
}
