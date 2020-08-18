package httpcli

import (
	"fmt"
	"testing"
	"time"
)

func TestReq(t *testing.T) {

	ts := time.Now().UnixNano() / 1000 / 1000
	tt := time.Now().Unix()
	time.Sleep(time.Second)
	ttt := time.Now().Unix()
	resp, err := NewBuilder("http://127.0.0.1:8000/api/v1/srv/users/online/update").
		Method(POST).
		Timeout(time.Second).
		SetContentType("application/json").
		MapBody(map[string]interface{}{
			"type": "jbzjh",
			"name": "127.0.0.1:8889",
			"cnt":  100,
			"ts":   ts,
		}).
		Do()

	fmt.Println(resp, err, tt, ttt)
}
