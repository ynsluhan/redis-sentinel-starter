package main

import (
	"github.com/spf13/viper"
	"context"
	. "github.com/ynsluhan/redis-sentinel-starter"
	"log"
	"os"
	"path"
)

/**
* @Author: yNsLuHan
* @Description:
* @File: main
* @Version: 1.0.0
* @Date: 2021/8/23 4:50 下午
 */
func main() {
	// 获取配置文件路径
	basePath, err := os.Getwd()
	var configPath = path.Join(basePath, "example")
	// 读取配置文件
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("application")
	config.AddConfigPath(configPath)
	// 初始化redis sentinel
	//
	err = config.ReadInConfig()
	//
	if err != nil {
		log.Fatal(err)
	}
	// 初始化redis sentinel
	InitRedisSentinel(config, "redis.sentinel")
	// 使用
	// 获取sentinel map
	sentinel := GetSentinel()
	// 获取node
	client := sentinel["node1"]
	// 查询
	result, _ := client.Get(context.Background(), "test").Result()
	//
	log.Println(result)
}
