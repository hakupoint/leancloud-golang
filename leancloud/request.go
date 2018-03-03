package leancloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type errorCode struct {
}

type Response struct {
	Code         float64   `json:"code"`
	Error        string    `json:"error"`
	CreatedAt    time.Time `json:"createdAt"`
	ObjectId     string    `json:"objectId"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Content      string    `json:"content"`
	PubUser      string    `json:"pubUser"`
	PubTimestamp float64   `json:"pubTimestamp"`
}

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

func fetch(r *http.Request, l *LeanCloud) (interface{}, error) {
	var respdata Response
	client := &http.Client{}
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &respdata)
	if err != nil {
		return nil, err
	}
	return respdata, nil
}

func (l *LeanCloud) createRequest(method, module string, arg interface{}) *http.Request {
	u, _ := url.Parse(API_URL + module)
	r := &http.Request{
		Method: method,
		URL:    u,
		Header: http.Header{},
	}
	switch t := arg.(type) {
	case url.Values:
		r.Form = t
	case string:
		var bf = ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("%+v", arg.(string))))
		r.Body = bf
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
