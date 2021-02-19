package router
import (
    "github.com/gin-gonic/gin"
    "fs-go-web/api/v1"
)
func Http(router *gin.Engine)  {
    apiRouter := router.Group("/api/v1/")
    {
        apiRouter.Any("/hello", v1.Hello)
        apiRouter.Any("/test_redis", v1.TestRedis)
    }
}
