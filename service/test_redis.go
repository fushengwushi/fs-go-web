package service

import (
	"fs-go-web/library"
	"fs-go-web/library/base"
	"fs-go-web/library/log"
	"github.com/gin-gonic/gin"
	"time"
)

const lastReqKey = "last:req:time"

func TestRedis(ctx *gin.Context, responseBody *library.ResponseBody)  {
	text := "上次请求时间为: "

	redisClient := base.GetNewRedisClient(ctx)

	val, err := redisClient.Get(lastReqKey)
	if err != nil {
		log.Error(map[string]interface{}{"get from redis error": err.Error()})
	}
	text += val

	err = redisClient.Set(lastReqKey, time.Now().Unix(), 600 * time.Second)
	if err != nil {
		log.Error(map[string]interface{}{"set redis error": err.Error()})
	}

	responseBody.SetData(text)
	return
}