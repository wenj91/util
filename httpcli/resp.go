package httpcli

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func ReadAll(resp *http.Response) ([]byte, error)  {
	if nil == resp {
		return nil, errors.New("resp is nil")
	}

	bs , err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return bs, nil
}