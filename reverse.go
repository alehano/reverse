// Copyright © 2015 Alexey Khalyapin - halyapin@gmail.com
//
// This program or package and any associated files are licensed under the
// GNU GENERAL PUBLIC LICENSE Version 2 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.gnu.org/licenses/gpl-2.0.html

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

// Adds a Url to the Store
func (us *urlStore) Add(urlName string, urlAddr string, params ...string) error {
	if _, ok := us.store[urlName]; ok {
		return errors.New("Url already exists")
	}

	tmpUrl := url{urlAddr, params}
	us.store[urlName] = tmpUrl
	return nil
}

// Adds a Url and panics if error
func (us urlStore) MustAdd(urlName string, urlAddr string, params ...string) string {
	err := us.Add(urlName, urlAddr, params...)
	if err != nil {
		panic(err)
	}
	return us.Get(urlName)
}

// Gets raw url string
func (us urlStore) Get(urlName string) string {
	return us.store[urlName].url
}

// Gets reversed url
func (us urlStore) Reverse(urlName string, params ...string) (string, error) {
	if len(params) != len(us.store[urlName].params) {
		return "", errors.New("Bad Url Reverse: mismatch params")
	}
	res := us.store[urlName].url
	for i, val := range params {
		res = strings.Replace(res, us.store[urlName].params[i], val, 1)
	}
	return res, nil
}

// Gets reversed url and panics if error
func (us urlStore) MustReverse(urlName string, params ...string) string {
	res, err := us.Reverse(urlName, params...)
	if err != nil {
		panic(err)
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
