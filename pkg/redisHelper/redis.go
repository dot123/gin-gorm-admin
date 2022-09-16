package redisHelper

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

// Set a key/value
func Set(ring *redis.Ring, key string, data interface{}, expiration time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = ring.Set(key, value, expiration).Err(); err != nil {
		return err
	}

	return nil
}

// Get a key
func Get(ring *redis.Ring, key string, out interface{}) (err error) {
	reply, err := ring.Get(key).Result()
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(reply), out); err != nil {
		return err
	}

	return nil
}

// LikeDeletes batch delete
func LikeDeletes(ring *redis.Ring, key string) error {
	val, err := ring.Do("KEYS", "*"+key+"*").Result()
	if err != nil {
		return err
	}
	keys := val.([]interface{})

	for _, value := range keys {
		k := value.(string)
		if err = ring.Del(k).Err(); err != nil {
			return err
		}
	}
	return nil
}

// Check ...
func Check(ring *redis.Ring, key string) (bool, error) {
	cmd := ring.Exists(key)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}
