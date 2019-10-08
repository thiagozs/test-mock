package dbs

import (
	"log"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

// Dispatch of dbs
type Dispatch struct {
	Redis *redis.Client
	Cache *cache.Codec
}

//RedisObject string
type RedisObject struct {
	Message []byte
}

// NewTestRedis mock
func NewTestRedis() (*Dispatch, error) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	codec := &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	_, err = client.Ping().Result()
	if err != nil {
		return &Dispatch{}, err

	}

	return &Dispatch{client, codec}, nil

}

// NewRedis connection with redis key value
func NewRedis() (*Dispatch, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	codec := &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	_, err := client.Ping().Result()
	if err != nil {
		return &Dispatch{}, err
	}

	return &Dispatch{client, codec}, nil
}

// Ping ping server
func (d *Dispatch) Ping() error {
	_, err := d.Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//QueueSet set item queue
func (d *Dispatch) QueueSet(queue string, json string) error {
	err := d.Redis.RPush(queue, json).Err()
	if err != nil {
		log.Printf("[Redis] RPUSH Error: %s\n", err)
		return err
	}
	return nil
}

// QueueGet get queue item
func (d *Dispatch) QueueGet(queue string) (string, error) {
	result, err := d.Redis.LPop(queue).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// QueueRangeList get queue item
func (d *Dispatch) QueueRangeList(queue string, start int64, end int64) ([]string, error) {
	result, err := d.Redis.LRange(queue, start, end).Result()
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

// QueueTrim remove item from queue
func (d *Dispatch) QueueTrim(queue string, start int64, end int64) (string, error) {
	result, err := d.Redis.LTrim(queue, start, end).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// QueueGetList list of item
func (d *Dispatch) QueueGetList(queue string, amount int) ([]string, error) {
	result := []string{}
	for i := 0; i < amount; i++ {
		data, err := d.QueueGet(queue)
		if err != nil {
			return result, err
		}
		result = append(result, data)
	}
	return result, nil
}

// QueueSize length of queue
func (d *Dispatch) QueueSize(queue string) (int, error) {
	result, err := d.Redis.LLen(queue).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Incr key
func (d *Dispatch) Incr(key string) (int, error) {
	result, err := d.Redis.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Decr key
func (d *Dispatch) Decr(key string) (int, error) {
	result, err := d.Redis.Decr(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Del key
func (d *Dispatch) Del(key string) (int, error) {
	result, err := d.Redis.Del(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Get key
func (d *Dispatch) Get(key string) (string, error) {
	result, err := d.Redis.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
