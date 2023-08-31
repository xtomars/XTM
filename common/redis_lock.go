/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: redis_lock.go
 * @Date: 2020/07/02
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Lock define
type Lock struct {
	resource string
	token    string
	conn     *redis.ClusterClient
	timeout  int
}

func (lock *Lock) tryLock() bool {
	return lock.conn.SetNX(lock.key(), lock.token, time.Duration(lock.timeout)*time.Second).Val()
}

// Unlock key
func (lock *Lock) Unlock() error {
	_, err := lock.conn.Del(lock.key()).Result()
	return err
}

func (lock *Lock) key() string {
	return fmt.Sprintf("redislock:%s", lock.resource)
}

// TryLock try to lock
func TryLock(token string, DefaulTimeout int) (lock *Lock, ok bool) {
	return TryLockWithTimeout(Redis.Client, notifyFailedLock, token, DefaulTimeout)
}

// TryLockWithTimeout try to lock with timeout
func TryLockWithTimeout(conn *redis.ClusterClient, resource string, token string, timeout int) (lock *Lock, ok bool) {
	lock = &Lock{resource, token, conn, timeout}

	ok = lock.tryLock()
	if !ok {
		lock = nil
	}

	return
}
