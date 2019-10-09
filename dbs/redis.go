package dbs

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

type DispatchRepository interface {
	Ping() error
	QueueSet(queue, json string) error
	QueueGet(queue string) (string, error)
	QueueRangeList(queue string, start, end int64) ([]string, error)
	QueueTrim(queue string, start, end int64) (string, error)
	QueueGetList(queue string, amount int) ([]string, error)
	QueueSize(queue string) (int, error)
	Incr(key string) (int, error)
	Decr(key string) (int, error)
	Del(key string) (int, error)
	Get(key string) (string, error)
}

type dispatchRepository struct {
	Redis *redis.Client
	Cache *cache.Codec
}

//RedisObject string
type RedisObject struct {
	Message []byte
}

// NewTestRedis mock
func NewRedis() DispatchRepository {
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

	return &dispatchRepository{client, codec}
}

// Ping ping server
func (d *dispatchRepository) Ping() error {
	_, err := d.Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//QueueSet set item queue
func (d *dispatchRepository) QueueSet(queue string, json string) error {
	err := d.Redis.RPush(queue, json).Err()
	if err != nil {
		return err
	}
	return nil
}

// QueueGet get queue item
func (d *dispatchRepository) QueueGet(queue string) (string, error) {
	result, err := d.Redis.LPop(queue).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// QueueRangeList get queue item
func (d *dispatchRepository) QueueRangeList(queue string, start, end int64) ([]string, error) {
	result, err := d.Redis.LRange(queue, start, end).Result()
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

// QueueTrim remove item from queue
func (d *dispatchRepository) QueueTrim(queue string, start, end int64) (string, error) {
	result, err := d.Redis.LTrim(queue, start, end).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

// QueueGetList list of item
func (d *dispatchRepository) QueueGetList(queue string, amount int) ([]string, error) {
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
func (d *dispatchRepository) QueueSize(queue string) (int, error) {
	result, err := d.Redis.LLen(queue).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Incr key
func (d *dispatchRepository) Incr(key string) (int, error) {
	result, err := d.Redis.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Decr key
func (d *dispatchRepository) Decr(key string) (int, error) {
	result, err := d.Redis.Decr(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Del key
func (d *dispatchRepository) Del(key string) (int, error) {
	result, err := d.Redis.Del(key).Result()
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

// Get key
func (d *dispatchRepository) Get(key string) (string, error) {
	result, err := d.Redis.Get(key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
