/**
单例模式注入连接对象后使用
redis base cmd packet
author Bill
*/
package hsredis

import (
	"demo/utils/json"
	"reflect"

	"github.com/garyburd/redigo/redis"
	"github.com/gpmgo/gopm/modules/log"
)

var (
	_r  *Redis
	con redis.Conn
)

type Redis struct {
}

func GetInstance() *Redis {
	if _r == nil {
		_r = &Redis{}
	}
	return _r
}

func (r *Redis) SetCon(connName string) {
	con = PoolConn(connName)
}

func (r *Redis) Set(key string, value interface{}, exTime int) bool {
	_value := value
	if reflect.TypeOf(value) == reflect.TypeOf(map[string]interface{}{}) {
		_value, _ = json.ToJsonString(value)
	}
	if exTime != 0 {
		_, err := con.Do("SET", key, _value, "Ex", exTime)
		if err != nil {
			log.Error("RedisErr : %s", err)
		}
		return true
	}
	_, err := con.Do("SET", key, _value)
	if err != nil {
		log.Error("RedisErr : %s", err)
	}
	return true
}

func (r *Redis) Del(key string) bool {
	_, err := con.Do("DEL", key)
	if err != nil {
		log.Error("RedisErr : %s", err)
	}
	return true
}

func (r *Redis) Get(key string) string {
	_reply, _ := redis.String(con.Do("GET", key))
	return _reply
}

func (r *Redis) GetInt(key string) int {
	_reply, _ := redis.Int(con.Do("GET", key))
	return _reply
}

func (r *Redis) GetJsonMap(key string) map[string]interface{} {
	_reply, _ := redis.String(con.Do("GET", key))
	return json.StringToJson(_reply)
}

func (r *Redis) HSet(table string, key string, value interface{}) bool {
	_value, _ := json.ToJsonString(value)
	_, err := con.Do("HSET", table, key, _value)
	if err != nil {
		log.Error("RedisErr : %s", err)
	}
	return true
}

func (r *Redis) SetNx(key string, val string) bool {
	res, err := con.Do("SETNX", key, val)
	if err != nil {
		log.Error("RedisErr : %s", err)
		return false
	}
	if res == "0" {
		return false
	}
	return true
}

func (r *Redis) Expire(key string, expireTime int) {
	con.Do("EXPIRE", key, expireTime)
}

func (r *Redis) Close() {
	con.Close()
}
