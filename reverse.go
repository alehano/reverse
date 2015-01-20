package reverse

import (
	"errors"
	"fmt"
	"strings"
)

var Urls *urlStore

func init() {
	Urls = &urlStore{store: make(map[string]url)}
}

type url struct {
	url    string
	params []string
}

type urlStore struct {
	store map[string]url
}

func (us *urlStore) Add(urlName string, urlAddr string, params ...string) error {
	if _, ok := us.store[urlName]; ok {
		errors.New("Url already exists")
	}

	tmpUrl := url{urlAddr, params}
	us.store[urlName] = tmpUrl
}

func (us urlStore) Get(urlName string) string {
	return us.store[urlName].url
}

func (us urlStore) Reverse(urlName string, params ...string) string, error {
	if len(params) != len(us.store[urlName].params) {
		errors.New("Bad Url Reverse: mismatch params")
	}
	res := us.store[urlName].url
	for i, val := range params {
		res = strings.Replace(res, us.store[urlName].params[i], val, 1)
	}
	return res
}

func (us urlStore) Sting() string {
	return fmt.Sprint(us.store)
}

// For testing
func (us urlStore) getParam(urlName string, num int) string {
	return us.store[urlName].params[num]
}
