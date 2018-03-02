package leancloud

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (l *LeanCloud) put(module string, data interface{}) *http.Request {
	r := l.createRequest(http.MethodPut, module, data)
	return r
}

func (l *LeanCloud) post(module string, data interface{}) *http.Request {
	r := l.createRequest(http.MethodPost, module, data)
	return r
}

func (l *LeanCloud) get(module string, data interface{}) *http.Request {
	r := l.createRequest(http.MethodGet, module, data)
	return r
}

func (l *LeanCloud) delete(module string, data interface{}) *http.Request {
	r := l.createRequest(http.MethodDelete, module, data)
	return r
}

func fetch(r *http.Request, l *LeanCloud) {
	client := &http.Client{}
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {

	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(b))
}

func (l *LeanCloud) createRequest(method, module string, data interface{}) *http.Request {
	var bf = ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("%+v", data)))
	u, _ := url.Parse(API_URL + module)
	r := &http.Request{
		Method: method,
		URL:    u,
		Body:   bf,
		Header: http.Header{},
	}
	l.sign()
	if l.sign_mode == SIGN_APP_KEY {
		r.Header.Set("X-LC-Sign", fmt.Sprintf("%s,%d", l.Sign, l.timestamp))
	}
	if l.sign_mode == SIGN_MASTER_KEY {
		r.Header.Set("X-LC-Sign", fmt.Sprintf("%s,%d,%s", l.Sign, l.timestamp, "master"))
	}
	r.Header.Set("X-LC-Id", l.Id)
	r.Header.Set("Content-Type", "application/json")
	return r
}
