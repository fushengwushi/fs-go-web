package base

import (
	"fs-go-web/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Ctx *gin.Context
}

func GetNewRedisClient(ctx *gin.Context) RedisClient {
	RedisConf := conf.BaseConf.Redis
	return RedisClient{Client: redis.NewClient(&redis.Options{
		Addr: RedisConf.Addr,
		Password: RedisConf.Password,
		PoolSize: RedisConf.PoolSize,
		DialTimeout: RedisConf.DialTimeout,
		ReadTimeout: RedisConf.ReadTimeout,
		WriteTimeout: RedisConf.WriteTimeout,
	}), Ctx: ctx}
}

func (r RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r RedisClient) Del(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r RedisClient) SetNX(key string, value interface{}, expiration time.Duration) error {
	return r.Client.SetNX(r.Ctx, key, value, expiration).Err()
}

func  (r RedisClient) Expire(key string, expiration time.Duration) error {
	return r.Client.Expire(r.Ctx, key, expiration).Err()
}

func  (r RedisClient) Exists(key string) error {
	return r.Client.Exists(r.Ctx, key).Err()
}

func (r RedisClient) HSet(key string, values ...interface{}) error {
	return r.Client.HSet(r.Ctx, key, values ...).Err()
}

func  (r RedisClient) HGet(key, field string) (string, error) {
	return r.Client.HGet(r.Ctx, key, field).Result()
}

func (r RedisClient) MSet(values ...interface{}) error {
	return r.Client.MSet(r.Ctx, values ...).Err()
}

func (r RedisClient) MGet(keys ...string) ([]interface{}, error) {
	return r.Client.MGet(r.Ctx, keys ...).Result()
}

func (r RedisClient) HSetNX(key, field string, value interface{}) error {
	return r.Client.HSetNX(r.Ctx, key, field, value).Err()
}

func  (r RedisClient) HExists(key, field string) error {
	return r.Client.HExists(r.Ctx, key, field).Err()
}

func  (r RedisClient) HDel(key, field string) error {
	return r.Client.HDel(r.Ctx, key, field).Err()
}

func (r RedisClient) HMSet(key string, values ...interface{}) error {
	return r.Client.HMSet(r.Ctx, key, values ...).Err()
}

func (r RedisClient) HMGet(key string, fields ...string) ([]interface{}, error) {
	return r.Client.HMGet(r.Ctx, key, fields ...).Result()
}

func  (r RedisClient) IncrBy(key string, value int64) error {
	return r.Client.IncrBy(r.Ctx, key, value).Err()
}

func  (r RedisClient) DecrBy(key string, value int64) error {
	return r.Client.DecrBy(r.Ctx, key, value).Err()
}

func  (r RedisClient) SAdd(key string, members ...interface{}) error {
	return r.Client.SAdd(r.Ctx, key, members ...).Err()
}

func  (r RedisClient) ZRange(key string, start, stop int64) ([]string, error) {
	return r.Client.ZRange(r.Ctx, key, start, stop).Result()
}

func  (r RedisClient) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.Client.ZRangeByScore(r.Ctx, key, opt).Result()
}

func  (r RedisClient) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return r.Client.ZRangeByScoreWithScores(r.Ctx, key, opt).Result()
}

func  (r RedisClient) ZRevRange(key string, start, stop int64) ([]string, error) {
	return r.Client.ZRevRange(r.Ctx, key, start, stop).Result()
}

func  (r RedisClient) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.Client.ZRevRangeByScore(r.Ctx, key, opt).Result()
}

func  (r RedisClient) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return r.Client.ZRevRangeByScoreWithScores(r.Ctx, key, opt).Result()
}

func  (r RedisClient) ZCard(key string) (int64, error) {
	return r.Client.ZCard(r.Ctx, key).Result()
}

func  (r RedisClient) SIsMember(key string, member interface{}) (bool, error) {
	return r.Client.SIsMember(r.Ctx, key, member).Result()
}

func  (r RedisClient) LPush(key string, values ...interface{}) error {
	return r.Client.LPush(r.Ctx, key, values ...).Err()
}

func  (r RedisClient) RPop(key string) (string, error) {
	return r.Client.RPop(r.Ctx, key).Result()
}

func  (r RedisClient) RPush(key string, values ...interface{}) error {
	return r.Client.LPush(r.Ctx, key, values ...).Err()
}

func  (r RedisClient) LPop(key string) (string, error) {
	return r.Client.RPop(r.Ctx, key).Result()
}

func  (r RedisClient) LRange(key string, start, stop int64) error {
	return r.Client.LRange(r.Ctx, key, start, stop).Err()
}