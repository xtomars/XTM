/*
 * *******************************************************************
 * @项目名称: redis
 * @文件名称: cluster.go
 * @Date: 2020/06/03
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package redis

import (
	"github.com/go-redis/redis"
)

// Cluster impl
type Cluster struct {
	Client *redis.ClusterClient
}

// NewCluster return the Redis
func NewCluster(addr, pass string, poolSize int) (*Cluster, error) {
	// connet redis server
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{addr},
		Password: pass,
		PoolSize: poolSize,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Errorf("Redis ping err:%v", err)
		return nil, err
	}
	return &Cluster{client}, nil
}

// AddToSet add a value to  set
func (r *Cluster) AddToSet(key, value string) error {
	if _, err := r.Client.SAdd(key, value).Result(); err != nil {
		log.Error("failed to add set")
		return err
	}
	return nil
}

// IsSetMember judge a value exist in set
func (r *Cluster) IsSetMember(key, value string) (bool, error) {
	return r.Client.SIsMember(key, value).Result()
}

// RemoveFromSet remove a value from  set
func (r *Cluster) RemoveFromSet(key, value string) error {
	if _, err := r.Client.SRem(key, value).Result(); err != nil {
		log.Error("failed remove value from set")
		return err
	}
	return nil
}

// GetSet get all values from set
func (r *Cluster) GetSet(key string) ([]string, error) {
	txs, err := r.Client.SMembers(key).Result()
	if err != nil {
		log.Error("failed get values from set")
		return nil, err
	}
	return txs, nil
}

// AddToList add value to list
func (r *Cluster) AddToList(key, vaule string) error {
	if _, err := r.Client.RPush(key, vaule).Result(); err != nil {
		log.Error("failed add vaule to redis list")
		return err
	}
	return nil
}

// GetList get values from list
func (r *Cluster) GetList(key string) ([]string, error) {
	// get list length
	length, err := r.Client.LLen(key).Result()
	if err != nil {
		log.Error("failed get list length")
		return nil, err
	}

	// get all values
	values, err := r.Client.LRange(key, 0, length).Result()
	if err != nil {
		log.Error("failed get values from redis list")
		return values, err
	}

	// remove all value from list
	for _, value := range values {
		r.Client.LRem(key, 0, value)
	}

	return values, nil
}

// Close the connection
func (r *Cluster) Close() {
	r.Client.Close()
}
