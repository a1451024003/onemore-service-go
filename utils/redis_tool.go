package utils

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

/*
获取值
    r := utils.Redix.Get()
	defer r.Close()

	var err error
	v, err := redis.String(r.Do("GET", "pool"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

设置值
	r := utils.Redix.Get()
	defer r.Close()

	var err error
	v, err := r.Do("SET", "pool", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
*/

//redis连接池
var (
	Redix *redis.Pool
)

/*
*addr: ip:port
*pwd: 密码，没有填""
*maxActive: 最大连接数
*idle: 连接空闲时长
 */
func InitRedix(addr, pwd string, maxActive, idle int) error {
	Redix = newPool(addr, pwd, maxActive, idle)
	r := Redix.Get()
	_, err := r.Do("PING")
	if err != nil {
		return err
	}
	r.Close()
	return err
}

//newPool
func newPool(server, password string, maxActive, idle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   maxActive,
		IdleTimeout: time.Second * time.Duration(idle),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
	}
}

func usecase() {
	//从redis连接池中获取一个连接
	r := Redix.Get()
	//把连接还给连接池
	defer r.Close()
	//TODO do your business
}
