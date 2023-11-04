// Package redis contains functions for using redis
package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"sports-day/conf"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	rdb redis.UniversalClient
)

// init needs to be called first to set up the rdb and rdbReplica values
func init() {
	options := redis.Options{
		Addr: conf.RedisConf.Addr,
	}
	replicaOptions := redis.Options{}
	if conf.RedisConf.EnableSSL {
		options.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
		replicaOptions.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	rdb = redis.NewClient(&options)
}

// Set sets a string value with given ttl against a key
// 0 ttl means no expiry
func Set(key string, value string, ttl time.Duration) error {
	if ttl < 0 {
		ttl = 0
	}
	if key == "" {
		return ErrorEmptyKey
	}
	key += suffix
	err := rdb.Set(context.Background(), key, value, ttl)
	if err != nil {
		return err.Err()
	}
	return nil
}

// SetStruct sets a struct object with given ttl against a key
func SetStruct(key string, obj interface{}, ttl time.Duration) error {
	valueBytes, err := json.Marshal(obj)
	if err != nil {
		return ErrorUnsupportedValue
	}
	return Set(key, string(valueBytes), ttl)
}

// Get returns value and error if any
// In case, key is not found it returns redis.Nil as error
func Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", ErrorEmptyKey
	}
	key += suffix
	ctx, cancelCtx := context.WithTimeout(ctx, contextTimeout)
	defer cancelCtx()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Delete deletes the key from redis
// It does not return an error if key is not found
func Delete(ctx context.Context, key string) error {
	if key == "" {
		return ErrorEmptyKey
	}
	key += suffix
	ctx, cancelCtx := context.WithTimeout(ctx, contextTimeout)
	defer cancelCtx()
	_, err := rdb.Del(ctx, key).Result()
	return err
}
