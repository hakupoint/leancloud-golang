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
type ScanResponse struct {
	Results []Response `json:"results"`
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

func fetch(r *http.Request, l *LeanCloud) (Response, error) {
	var respdata Response
	resp, err := fetchDo(r, respdata)
	if err != nil {
		return respdata, err
	}
	json.Unmarshal(resp, &respdata)
	return respdata, nil
}

func fetchList(r *http.Request, l *LeanCloud) (ScanResponse, error) {
	var respList ScanResponse
	resp, err := fetchDo(r, respList)
	if err != nil {
		return respList, err
	}
	json.Unmarshal(resp, &respList)
	return respList, nil
}
func fetchDo(r *http.Request, inter interface{}) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return b, nil
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
