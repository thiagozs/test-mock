package dbs_test

/*
// MockRedisClient fnc
type MockRedisClient interface {
	Ping() error
	QueueSet(queue string, json string) error
	QueueGet(queue string) (string, error)
	QueueRangeList(queue string, start int64, end int64) ([]string, error)
	QueueTrim(queue string, start int64, end int64) (string, error)
	QueueGetList(queue string, amount int) ([]string, error)
	QueueSize(queue string) (int, error)
	Incr(key string) (int, error)
	Decr(key string) (int, error)
	Del(key string) (int, error)
	Get(key string) (string, error)
}

type mockRedis struct{}

func (m *mockRedis) Ping() error {
	if m.PingFunc != nil {
		return m.PingFunc()
	}
	return fmt.Errorf("Ping Error")
}

func (m *mockRedis) QueueSet(queue string, json string) error {
	if m.QueueSetFunc != nil {
		return m.QueueSetFunc(queue, json)
	}
	return fmt.Errorf("QueueSet %s Error", queue)
}

func (m *mockRedis) QueueGet(queue string) (string, error) {
	if m.QueueGetFunc != nil {
		return m.QueueGetFunc(queue)
	}
	return "", fmt.Errorf("QueueGet %s Error", queue)
}

func (m *mockRedis) QueueRangeList(queue string, start, end int64) ([]string, error) {
	if m.QueueRangeListFunc != nil {
		return m.QueueRangeListFunc(queue, start, end)
	}
	return []string{}, fmt.Errorf("QueueRangeList %s %d %d Error", queue, start, end)
}

func (m *mockRedis) QueueTrim(queue string, start, end int64) (string, error) {
	if m.QueueTrimFunc != nil {
		return m.QueueTrimFunc(queue, start, end)
	}
	return "", fmt.Errorf("QueueTrim %s %d %d Error", queue, start, end)
}

func (m *mockRedis) QueueGetList(queue string, amount int) ([]string, error) {
	if m.QueueGetListFunc != nil {
		return m.QueueGetListFunc(queue, amount)
	}
	return []string{}, fmt.Errorf("QueueGetList %s %d Error", queue, amount)
}

func (m *mockRedis) QueueSize(queue string) (int, error) {
	if m.QueueSizeFunc != nil {
		return m.QueueSizeFunc(queue)
	}
	return 0, fmt.Errorf("QueueSize %s Error", queue)
}

func (m *mockRedis) Incr(key string) (int, error) {
	if m.IncrFunc != nil {
		return m.IncrFunc(key)
	}
	return 0, fmt.Errorf("Incr %s Error", key)
}

func (m *mockRedis) Decr(key string) (int, error) {
	if m.DecrFunc != nil {
		return m.DecrFunc(key)
	}
	return 0, fmt.Errorf("Decr %s Error", key)
}

func (m *mockRedis) Del(key string) (int, error) {
	if m.DelFunc != nil {
		return m.DelFunc(key)
	}
	return 0, fmt.Errorf("Del %s Error", key)
}

func (m *mockRedis) Get(key string) (string, error) {
	if m.GetFunc != nil {
		//return key, nil
		return m.GetFunc(key)
	}
	return "", fmt.Errorf("Get %s Error", key)
}
*/
