package hsredis

import (
	"demo/config"
	"github.com/garyburd/redigo/redis"
	"time"
	"net/url"
	"strconv"
	"fmt"
	"sync"
)

var (
	client      map[string]redis.Conn
	clientOne   redis.Conn
	poolClient  map[string]*redis.Pool
	RedisClient *redis.Pool
	LockSign    sync.Mutex
)

/**
	建立连接池管理机制(redis无法长连接)
 */
func PoolConn(name string) redis.Conn {
	if poolClient[name] != nil {
		return poolClient[name].Get()
	}
	poolClient = map[string]*redis.Pool{}
	config, _ := config.GetENVConfigs().String("redis." + name)
	u, _ := url.Parse(config)
	host := u.Host
	params, _ := url.ParseQuery(u.RawQuery)
	db := 0
	password := ""
	if params["db"] != nil {
		db, _ = strconv.Atoi(params["db"][0])
	}
	if params["password"] != nil {
		password = params["password"][0]
	}
	poolClient[name] = &redis.Pool{
		MaxIdle:     300, //TODO::配置化,空闲等待个数
		MaxActive:   300, //TODO::配置化,最大激活连接数
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}

			if password != ""{
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if db > 0{
				c.Do("SELECT", db)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return poolClient[name].Get()
}
//请使用这种连接
func Conn(name string) redis.Conn {
	LockSign.Lock()
	config, _ := config.GetENVConfigs().String("redis." + name)
	u, _ := url.Parse(config)
	host := u.Host
	params, _ := url.ParseQuery(u.RawQuery)
	db := 0
	var password string
	if params["db"] != nil {
		db, _ = strconv.Atoi(params["db"][0])
	}
	if params["password"] != nil {
		password = params["password"][0]
	}
	clientOne, _ = redis.Dial("tcp", host)
	if password != "" {
		if _, err := clientOne.Do("AUTH", password); err != nil {
			client[name].Close()
			fmt.Print("AUTH Error!")
			return nil
		}
	}
	if db > 0 {
		if _, err3 := clientOne.Do("SELECT", db); err3 != nil {
			client[name].Close()
			fmt.Print("Select Db Error!")
			return nil
		}
	}
	LockSign.Unlock()
	return clientOne
}
