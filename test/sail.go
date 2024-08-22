package main

import (
	"fmt"
	"github.com/HYY-yu/sail-client"
	"time"
)

type Config struct {
	Name    string
	Host    string
	Port    string
	Mode    string
	UserRpc struct {
		Etcd struct {
			Hosts []string
			Key   string
		}
	}

	Database string

	Redisx struct {
		Host string
		Pass string
	}
	JwtAuth struct {
		AccessSecret string
	}
}

// main
//
//	@Description: 在sail配置中心上设置每个微服务的配置，通过etcd进行拿取，但是代码中要进行加载
func main() {
	var cfg Config

	s := sail.New(&sail.MetaConfig{
		ETCDEndpoints:  "10.3.237.104:3379",
		ProjectKey:     "98c6f2c2287f4c73cea3d40ae7ec3ff2", // sail项目管理的key
		Namespace:      "user",
		Configs:        "user-api.yaml",
		ConfigFilePath: "./conf",
		LogLevel:       "DEBUG",
	}, sail.WithOnConfigChange(func(configFileKey string, s *sail.Sail) { // 配置改变之后动态加载
		if s.Err() != nil {
			fmt.Println(s.Err())
			return
		}

		// 拼接的请求字符串
		fmt.Println(s.Pull())

		// 拿到配置
		v, err := s.MergeVipers()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 反序列化
		v.Unmarshal(&cfg)
		fmt.Println(cfg, "\n", cfg.Database)
	}))

	if s.Err() != nil {
		fmt.Println(s.Err())
		return
	}

	// 拼接的请求字符串
	fmt.Println(s.Pull())

	// 拿到配置
	v, err := s.MergeVipers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 反序列化
	v.Unmarshal(&cfg)
	fmt.Println(cfg, "\n", cfg.Database)

	for {
		time.Sleep(time.Second)
	}
}
