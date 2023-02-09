package util

import (
	"os"
	"strconv"
)

func GetEnvVal(key, defaultVal string) string {
	//os.Getenv() 将返回一个空字符串，使用 LookupEnv 来区分空值和未设置值。
	val, exist := os.LookupEnv(key)
	if !exist {
		return defaultVal
	}
	return val
}

func GetEnvIntVal(key string, defaultVal int) int {
	valStr, exist := os.LookupEnv(key)
	if !exist {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

func GetEnvBoolVal(key string, defaultVal bool) bool {
	valStr, exist := os.LookupEnv(key)
	if !exist {
		return defaultVal
	}
	val, err := strconv.ParseBool(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}
