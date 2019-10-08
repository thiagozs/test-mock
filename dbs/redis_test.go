package dbs_test

import (
	"errors"
	"testing"

	"github.com/thiagozs/test-mock/dbs"
)

var redis = dbs.Dispatch{}

func init() {
	rr, err := dbs.NewTestRedis()
	if err != nil {
		panic(err)
	}
	redis = *rr
}
func TestGetFunc(t *testing.T) {
	cases := []struct {
		redis dbs.Dispatch
		key   string
		err   error
	}{
		{redis, "aaa", nil},
		{redis, "", nil},
		{redis, "", errors.New("Fail get m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.Get(cc.key)
		switch i {
		case 0, 1:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if err == nil && cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestDelFunc(t *testing.T) {
	cases := []struct {
		redis  dbs.Dispatch
		item   string
		result int
		err    error
	}{
		{redis, "aaa", 1, nil},
		{redis, "", 0, nil},
		{redis, "", 0, errors.New("Fail del m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.Del(cc.item)
		switch i {
		case 0, 1:
			if cc.err != nil && (cc.result == 0 || cc.result == 1) {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestDecrFunc(t *testing.T) {
	cases := []struct {
		redis  dbs.Dispatch
		item   string
		result int
		err    error
	}{
		{redis, "bbb", 1, nil},
		{redis, "bbba", 1, nil},
		{redis, "bbbb", 1, errors.New("Fail decr m")},
		{redis, "bbb", 0, nil},
		{redis, "bbba", 0, nil},
		{redis, "bbbb", 0, errors.New("Fail decr m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.Decr(cc.item)
		switch i {
		case 0, 1, 3, 4:
			if cc.err != nil && (cc.result == 0 || cc.result == 1) {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestIncrFunc(t *testing.T) {
	cases := []struct {
		redis  dbs.Dispatch
		item   string
		result int
		err    error
	}{
		{redis, "bbb", 1, nil},
		{redis, "bbba", 1, nil},
		{redis, "bbbb", 1, errors.New("Fail incr m")},
		{redis, "bbb", 0, nil},
		{redis, "bbba", 0, nil},
		{redis, "bbbb", 0, errors.New("Fail incr m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.Incr(cc.item)
		switch i {
		case 0, 1, 3, 4:
			if cc.err != nil && (cc.result == 0 || cc.result == 1) {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestQueueSizeFunc(t *testing.T) {
	cases := []struct {
		redis  dbs.Dispatch
		item   string
		result int
		err    error
	}{
		{redis, "worker", 1, nil},
		{redis, "worker", 1, nil},
		{redis, "worker", 1, errors.New("Fail QueueSize m")},
		{redis, "worker", 0, nil},
		{redis, "worker", 0, nil},
		{redis, "worker", 0, errors.New("Fail QueueSize m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.QueueSize(cc.item)
		switch i {
		case 0, 1, 3, 4:
			if cc.err != nil && (cc.result == 0 || cc.result == 1) {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestQueueGetListFunc(t *testing.T) {
	cases := []struct {
		redis  dbs.Dispatch
		item   string
		amount int
		err    error
	}{
		{redis, "worker", 1, nil},
		{redis, "worker", 10, nil},
		{redis, "worker", 100, errors.New("Fail QueueGetList m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.QueueGetList(cc.item, cc.amount)
		switch i {
		case 0, 1, 3, 4:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestQueueTrimFunc(t *testing.T) {
	cases := []struct {
		redis dbs.Dispatch
		item  string
		start int64
		end   int64
		err   error
	}{
		{redis, "worker", 1, 3, nil},
		{redis, "worker", 10, 15, nil},
		{redis, "worker", 100, 150, errors.New("Fail QueueTrim m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.QueueTrim(cc.item, cc.start, cc.end)
		switch i {
		case 0, 1:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestQueueRangeListFunc(t *testing.T) {
	cases := []struct {
		redis dbs.Dispatch
		item  string
		start int64
		end   int64
		err   error
	}{
		{redis, "worker", 1, 3, nil},
		{redis, "worker", 10, 15, nil},
		{redis, "worker", 100, 150, errors.New("Fail QueueRangeList m")},
	}

	for i, cc := range cases {
		_, err := cc.redis.QueueRangeList(cc.item, cc.start, cc.end)
		switch i {
		case 0, 1:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestQueueSetFunc(t *testing.T) {
	cases := []struct {
		redis dbs.Dispatch
		item  string
		json  string
		err   error
	}{
		{redis, "worker", `{"name":"nameA"}`, nil},
		{redis, "worker", `{"name":"nameB"}`, nil},
		{redis, "worker", `{"name":"nameC"}`, errors.New("Fail QueueSet m")},
	}

	for i, cc := range cases {
		err := cc.redis.QueueSet(cc.item, cc.json)
		switch i {
		case 0, 1:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}

func TestPingFunc(t *testing.T) {
	cases := []struct {
		redis dbs.Dispatch
		err   error
	}{
		{redis, nil},
		{redis, errors.New("Fail Ping m")},
	}

	for i, cc := range cases {
		err := cc.redis.Ping()
		switch i {
		case 0:
			if err != nil && cc.err != nil {
				t.Errorf("Case %d, expected no erros, but got %v", i, err)
			}
		default:
			if cc.err == nil {
				t.Errorf("Case %d, expected erros, but got %v", i, err)
			}
		}
	}
}
