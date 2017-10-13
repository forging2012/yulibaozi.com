package cache

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

type RedisUtils struct {
}

// EXISTS key是否存在
func (m *RedisUtils) EXISTS(key string) (exists bool, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	exists, err = redis.Bool(rc.Do("EXISTS", key))
	return
}

// HMSET redis HMSET操作(HASH)
func (m *RedisUtils) HMSET(key string, obj interface{}) error {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := redis.String(rc.Do("HMSET", redis.Args{}.Add(key).AddFlat(obj)...))
	return err
}

// LPUSH 消息队列的写入操作
// key string 要写入的消息队列
// value interface{}写入的值
func (m *RedisUtils) LPUSH(key string, value interface{}) (err error) {
	rc := RedisClient.Get()
	var resultId int64
	defer rc.Close()
	resultId, err = redis.Int64(rc.Do("LPUSH", redis.Args{}.Add(key).Add(value)...))
	if resultId <= 0 {
		err = errors.New("插入出错")
	}
	return
}

// BRPOP 消息队列的取出数据操作(阻塞)
// key 操作的队列
// timeout 超时
func (m *RedisUtils) BRPOP(key string, timeout int64) (strings []string, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	var (
		values []interface{}
	)
	values, err = redis.Values(rc.Do("BRPOP", redis.Args{}.Add(key).AddFlat(timeout)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &strings)
	if err != nil {
		return
	}
	if len(strings) < 2 {
		err = errors.New("未取到任何值")
	} else {
		strings = strings[1:]
	}
	return
}

// HMSETWEIGHT 添加到HASH的时候同时添加到ZSET
// weight: int64 ZSET的权重值
// zset: ZSET的key
func (m *RedisUtils) HMSETWEIGHT(key, zset string, weight int64, obj interface{}) (err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	var result string
	result, err = redis.String(rc.Do("HMSET", redis.Args{}.Add(key).AddFlat(obj)...))
	if err != nil && result != "OK" {
		return
	}
	_, err = rc.Do("ZADD", redis.Args{}.Add(zset).Add(weight).Add(key)...)
	return
}

// ZADD 添加到ZSET
// key: string sorted set的key
// field: 字段名
// weight: 权重即score
func (m *RedisUtils) ZADD(key, field string, weight int64) (err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	_, err = redis.Int64(rc.Do("ZADD", redis.Args{}.Add(key).Add(weight).Add(field)...))
	return
}

// HDELWEIGHT redis 删除HASH 同时删除ZSET
// key: string hash/zset的fileds
// fields: 需要删除的hash fields
func (m *RedisUtils) HDELWEIGHT(key, zset string, fields interface{}) (err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	var resultInt int64
	resultInt, err = redis.Int64(rc.Do("HDEL", redis.Args{}.Add(key).AddFlat(fields)...))
	if err != nil && resultInt == 0 {
		if err != nil {
			return
		} else {
			err = errors.New("未删除redis任何数据")
			return
		}
	} else {
		resultInt, err = redis.Int64(rc.Do("ZREM", redis.Args{}.Add(zset).Add(key)...))
		if resultInt == 0 {
			err = errors.New("未删除redis任何数据2")
		}
		return
	}
}

// HGETALL redis HGETALL操作
func (m *RedisUtils) HGETALL(key, obj interface{}) (err error) {
	var v []interface{}
	rc := RedisClient.Get()
	defer rc.Close()
	v, err = redis.Values(rc.Do("HGETALL", key))
	if len(v) == 0 {
		err = errors.New("未查询到该数据")
		return
	}
	if err != nil {
		return
	}
	err = redis.ScanStruct(v, obj)
	return
}

// GetZsetFields 顺序获取分页区间的zset里面的Fields
// key 需要操作的zset
// offset 上区间
// limit 下区间
func (m *RedisUtils) GetZsetFields(key string, offset, limit int64) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := RedisClient.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZRANGE", redis.Args{}.Add(key).Add(offset).Add(limit)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New("未查找到数据")
		return
	}
	return
}

// GetZsetFields 倒序获取分页区间的zset里面的Fields
// key 需要操作的zset
// offset 上区间
// limit 下区间
func (m *RedisUtils) GetZsetREVFields(key string, offset, limit int64) (slice []string, err error) {
	var (
		values []interface{}
	)
	rc := RedisClient.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do("ZREVRANGE", redis.Args{}.Add(key).Add(offset).Add(limit)...))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New("未查找到数据")
		return
	}
	return
}

// CountByKey 统计指定zset的总记录数
// key string:指定统计那个zset
func (m *RedisUtils) CountByKey(key string) (count int64, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	count, err = redis.Int64(rc.Do("ZCARD", key))
	return
}

func (m *RedisUtils) Keys(key string) (slice []string, err error) {
	rc := RedisClient.Get()
	var (
		values []interface{}
	)
	defer rc.Close()
	values, err = redis.Values(rc.Do("keys", key))
	if err != nil {
		return
	}
	err = redis.ScanSlice(values, &slice)
	if err != nil {
		return
	}
	if len(slice) <= 0 {
		err = errors.New("未查找到数据")
		return
	}
	return
}

// SETEX 在redis中写入string
// 返回值：设置成功时返回OK；seconds无效时，返回错误
// key string redis key
// ex int  key的存活时间  单位s
// value string key对应的值
func (m *RedisUtils) SETEX(key string, ex int64, value string) (ok bool, err error) {
	var (
		result string
	)
	rc := RedisClient.Get()
	defer rc.Close()
	result, err = redis.String(rc.Do("SETEX", redis.Args{}.Add(key).Add(ex).Add(value)...))
	if err != nil && result != "OK" {

		return false, err
	}
	return true, err
}

// GET 从redis的string中取出值
//  返回值 get不会出现错误，当取到了值返回值，当没有取到就返回null，在这里，我们认为的当为null的时候就返回错误
func (m *RedisUtils) GET(key string) (value string, err error) {
	rc := RedisClient.Get()
	defer rc.Close()
	value, err = redis.String(rc.Do("GET", key))
	return
}

