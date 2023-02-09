package cache

import (
	"bytes"
	. "little_robot/db"
	util "little_robot/pkg/utils"
	"sync"
	"time"
)

type RedisHotReloadStorage struct {
	Key    string
	reader *bytes.Reader
	lock   sync.Mutex
}

func NewRedisHotReloadStorage(key string) *RedisHotReloadStorage {
	return &RedisHotReloadStorage{Key: key}
}

// Load 重写热登录数据加载，从Redis取数据
func (f *RedisHotReloadStorage) Read(p []byte) (n int, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.reader == nil {
		// 从Redis获取热登录数据
		data, err := GetRedisInstance().Get(f.Key)
		if err != nil {
			util.LogrusObj.Errorf("读取热登录数据失败: %v", err)
			return 0, err
		}
		f.reader = bytes.NewReader([]byte(data))
	}
	return f.reader.Read(p)
}

// Dump 重写更新热登录数据，保存到Redis
func (f *RedisHotReloadStorage) Write(p []byte) (n int, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	_, err = GetRedisInstance().Set(f.Key, string(p), 2*24*time.Hour)
	if err != nil {
		util.LogrusObj.Errorf("保存热登陆数据失败: %v", err.Error())
		return
	}
	return len(p), nil
}

// Close 需要关闭
func (f *RedisHotReloadStorage) Close() error {
	f.reader = nil
	return nil
}
