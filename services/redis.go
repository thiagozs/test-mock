package services

type DispatchServices interface {
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

type dispatchServices struct {
	service DispatchServices
}

func NewDispatchServices(d DispatchServices) DispatchServices {
	return &dispatchServices{
		d,
	}
}

// Ping ping server
func (d *dispatchServices) Ping() error {
	return d.service.Ping()
}

//QueueSet set item queue
func (d *dispatchServices) QueueSet(queue string, json string) error {
	return d.service.QueueSet(queue, json)
}

// QueueGet get queue item
func (d *dispatchServices) QueueGet(queue string) (string, error) {
	return d.service.QueueGet(queue)
}

// QueueRangeList get queue item
func (d *dispatchServices) QueueRangeList(queue string, start, end int64) ([]string, error) {
	return d.service.QueueRangeList(queue, start, end)
}

// QueueTrim remove item from queue
func (d *dispatchServices) QueueTrim(queue string, start, end int64) (string, error) {
	return d.service.QueueTrim(queue, start, end)
}

// QueueGetList list of item
func (d *dispatchServices) QueueGetList(queue string, amount int) ([]string, error) {
	return d.service.QueueGetList(queue, amount)
}

// QueueSize length of queue
func (d *dispatchServices) QueueSize(queue string) (int, error) {
	return d.service.QueueSize(queue)
}

// Incr key
func (d *dispatchServices) Incr(key string) (int, error) {
	return d.service.Incr(key)
}

// Decr key
func (d *dispatchServices) Decr(key string) (int, error) {
	return d.service.Decr(key)
}

// Del key
func (d *dispatchServices) Del(key string) (int, error) {
	return d.service.Del(key)
}

// Get key
func (d *dispatchServices) Get(key string) (string, error) {
	return d.service.Get(key)
}
