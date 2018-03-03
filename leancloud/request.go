package leancloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func put(l *LeanCloud, module string, data interface{}) *http.Request {
	return createRequest(l, http.MethodPut, module, data)
}

func post(l *LeanCloud, module string, data interface{}) *http.Request {
	return createRequest(l, http.MethodPost, module, data)
}

func get(l *LeanCloud, module string, data interface{}) *http.Request {
	return createRequest(l, http.MethodGet, module, data)
}

func delete(l *LeanCloud, module string, data interface{}) *http.Request {
	return createRequest(l, http.MethodDelete, module, data)
}

func fetch(r *http.Request, l *LeanCloud) (Response, error) {
	var respdata Response
	resp, err := fetchDo(r)
	if err != nil {
		return respdata, err
	}
	fmt.Printf("\n%+v", string(resp))
	json.Unmarshal(resp, &respdata)
	return respdata, nil
}

func fetchList(r *http.Request, l *LeanCloud) (ScanResponse, error) {
	var respList ScanResponse
	resp, err := fetchDo(r)
	if err != nil {
		return respList, err
	}
	json.Unmarshal(resp, &respList)
	return respList, nil
}
func fetchDo(r *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil
}

func createRequest(l *LeanCloud, method, module string, arg interface{}) *http.Request {
	u, _ := url.Parse(API_URL + module)
	r := &http.Request{
		Method: method,
		URL:    u,
		Header: http.Header{},
	}
	typefilter(r, arg)
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

func typefilter(r *http.Request, arg interface{}) {
	switch t := arg.(type) {
	case url.Values:
		r.Form = t
	case string:
		var bf = ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("%+v", t)))
		r.Body = bf
	case RequestData:
		b, _ := json.Marshal(t)
		r.Body = ioutil.NopCloser(bytes.NewBufferString(string(b)))
	}
}
