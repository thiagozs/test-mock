package dbs_test

import "fmt"

// MockRedisClient fnc
type MockRedisClient struct {
	PingFunc           func() error
	QueueSetFunc       func(queue string, json string) error
	QueueGetFunc       func(queue string) (string, error)
	QueueRangeListFunc func(queue string, start int64, end int64) ([]string, error)
	QueueTrimFunc      func(queue string, start int64, end int64) (string, error)
	QueueGetListFunc   func(queue string, amount int) ([]string, error)
	QueueSizeFunc      func(queue string) (int, error)
	IncrFunc           func(key string) (int, error)
	DecrFunc           func(key string) (int, error)
	DelFunc            func(key string) (int, error)
	GetFunc            func(key string) (string, error)
}

func (m *MockRedisClient) Ping() error {
	if m.PingFunc != nil {
		return m.PingFunc()
	}
	return fmt.Errorf("Ping Error")
}

func (m *MockRedisClient) QueueSet(queue string, json string) error {
	if m.QueueSetFunc != nil {
		return m.QueueSetFunc(queue, json)
	}
	return fmt.Errorf("QueueSet %s Error", queue)
}

func (m *MockRedisClient) QueueGet(queue string) (string, error) {
	if m.QueueGetFunc != nil {
		return m.QueueGetFunc(queue)
	}
	return "", fmt.Errorf("QueueGet %s Error", queue)
}

func (m *MockRedisClient) QueueRangeList(queue string, start, end int64) ([]string, error) {
	if m.QueueRangeListFunc != nil {
		return m.QueueRangeListFunc(queue, start, end)
	}
	return []string{}, fmt.Errorf("QueueRangeList %s %d %d Error", queue, start, end)
}

func (m *MockRedisClient) QueueTrim(queue string, start, end int64) (string, error) {
	if m.QueueTrimFunc != nil {
		return m.QueueTrimFunc(queue, start, end)
	}
	return "", fmt.Errorf("QueueTrim %s %d %d Error", queue, start, end)
}

func (m *MockRedisClient) QueueGetList(queue string, amount int) ([]string, error) {
	if m.QueueGetListFunc != nil {
		return m.QueueGetListFunc(queue, amount)
	}
	return []string{}, fmt.Errorf("QueueGetList %s %d Error", queue, amount)
}

func (m *MockRedisClient) QueueSize(queue string) (int, error) {
	if m.QueueSizeFunc != nil {
		return m.QueueSizeFunc(queue)
	}
	return 0, fmt.Errorf("QueueSize %s Error", queue)
}

func (m *MockRedisClient) Incr(key string) (int, error) {
	if m.IncrFunc != nil {
		return m.IncrFunc(key)
	}
	return 0, fmt.Errorf("Incr %s Error", key)
}

func (m *MockRedisClient) Decr(key string) (int, error) {
	if m.DecrFunc != nil {
		return m.DecrFunc(key)
	}
	return 0, fmt.Errorf("Decr %s Error", key)
}

func (m *MockRedisClient) Del(key string) (int, error) {
	if m.DelFunc != nil {
		return m.DelFunc(key)
	}
	return 0, fmt.Errorf("Del %s Error", key)
}

func (m *MockRedisClient) Get(key string) (string, error) {
	if m.GetFunc != nil {
		//return key, nil
		return m.GetFunc(key)
	}
	return "", fmt.Errorf("Get %s Error", key)
}
