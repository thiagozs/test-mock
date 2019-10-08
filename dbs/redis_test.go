package dbs

import (
	"errors"
	"fmt"
	"testing"
)

func TestWhatEver(t *testing.T) {
	cases := []struct {
		redisClient RedisClient
		item        string
		err         error
	}{
		{
			&MockRedisClient{},
			"aaa",
			errors.New("Get key Error"),
		},
	}

	for _, cc := range cases {
		_, err := cc.redisClient.Get(cc.item)
		fmt.Println(err.Error())
		if err != nil {
			t.Error(err)
		}
	}

}
