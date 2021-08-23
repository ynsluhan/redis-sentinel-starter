package Starter

import (
	"github.com/spf13/viper"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
)

//
var sentinelMap = make(map[string]*redis.Client, 1)

/**
* @Author: yNsLuHan
* @Description:
* @File: main
* @Version: 1.0.0
* @Date: 2021/8/23 3:31 下午
 */
//func InitRedisSentinel(configPath, configType, configName string) {
//	// 创建viper 读取配置文件
//	config := viper.New()
//	config.SetConfigName(configName)
//	config.SetConfigType(configType)
//	config.AddConfigPath(configPath)
//	//
//	err := config.ReadInConfig()
//	//
//	if err != nil {
//		log.Fatal(err)
//	}
//	//
//	get := config.Get("redis.sentinel")
//	//
//	//
//	for s, i := range get.(map[string]interface{}) {
//		SetRedisDb(i.(map[string]interface{}), s)
//	}
//}

func InitRedisSentinel(config *viper.Viper, name string) {
	get := config.Get(name)
	for s, i := range get.(map[string]interface{}) {
		SetRedisDb(i.(map[string]interface{}), s)
	}
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-06-08 15:26:02
 */
func SetRedisDb(m map[string]interface{}, nodeName string) {
	address := GetStringMustOption("address", m).(string)
	name := GetStringMustOption("name", m).(string)
	password := GetStringMustOption("password", m).(string)
	db := GetIntMustOption("db", m).(int)
	poolSize := GetIntMustOption("pool-size", m).(int)
	minIdleConn := GetIntMustOption("min-idle-conns", m).(int)
	// 将地址进行切割
	addressList := strings.Split(address, ",")
	//建立连接
	con := &redis.FailoverOptions{
		// master name.
		MasterName: name,
		// sentinel list
		SentinelAddrs: addressList,
		// 连接密码
		Password: password,
		// db
		DB: db,
		// 连接池个数
		PoolSize: poolSize,
		// 最小空闲个数
		MinIdleConns: minIdleConn,
	}
	client := redis.NewFailoverClient(con)
	//
	sentinelMap[nodeName] = client
	log.Println("INFO Redis sentinel node:", nodeName, "name:", name, "address:", address, "db:", db, "init success...")
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-08-23 16:07:48
 * @return map[string]*redis.Client
 */
func GetSentinel() map[string]*redis.Client {
	return sentinelMap
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-08-23 12:12:08
 * @param optionName
 * @param data
 * @return string
 */
func GetStringMustOption(optionName string, data map[string]interface{}) interface{} {
	h := data[optionName]

	if h == nil {
		log.Fatal("ERROR redis：", optionName, " 字段为空")
	}
	return h.(interface{})
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-08-23 12:12:08
 * @param optionName
 * @param data
 * @return string
 */
func GetStringOption(optionName string, data map[string]interface{}) interface{} {
	h := data[optionName]

	if h == nil {
		return nil
	}

	return h.(interface{})
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-08-23 12:12:08
 * @param optionName
 * @param data
 * @return string
 */
func GetIntOption(optionName string, data map[string]interface{}) interface{} {
	h := data[optionName]

	if h == nil {
		return nil
	}

	return h.(interface{})
}

/**
 * @Author yNsLuHan
 * @Description:
 * @Time 2021-08-23 12:12:08
 * @param optionName
 * @param data
 * @return string
 */
func GetIntMustOption(optionName string, data map[string]interface{}) interface{} {
	h := data[optionName]

	if h == nil {
		log.Fatal("ERROR redis：", optionName, " 字段为空")
	}

	return h.(interface{})
}
