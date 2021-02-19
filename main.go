package main
import (
    "fs-go-web/conf"
    "fs-go-web/library/middleware"
    "github.com/gin-gonic/gin"
    "fs-go-web/router"
)
func main()  {
    // 1.创建路由
    r := gin.Default()
    r.Use(middleware.LoggerToFile())

    conf.InitConf()

    // 2.绑定路由规则，执行的函数
    // gin.Context，封装了request和response
    router.Http(r)
    // 3.监听端口，默认在8080
    // Run("里面不指定端口号默认为8088")
    r.Run(":8088")
}
