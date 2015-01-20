package reverse

import (
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

// TODO: add errors
func (us *urlStore) Add(urlName string, urlAddr string, params ...string) {
	if _, ok := us.store[urlName]; ok {
		panic("Url already exists")
	}

	tmpUrl := url{urlAddr, params}
	us.store[urlName] = tmpUrl
}

// TODO: set out of package (with my ErrorStore)

// Set Name, Url, Parameters in url
func (us urlStore) Set(urlName string, urlAddr string, params ...string) string {
	us.Add(urlName, urlAddr, params...)
	return urlAddr
}

func (us urlStore) Get(urlName string) string {
	return us.store[urlName].url
}

// TODO: add errors
func (us urlStore) Reverse(urlName string, params ...string) string {
	if len(params) != len(us.store[urlName].params) {
		panic("Bad Url Reverse")
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
