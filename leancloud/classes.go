package leancloud

import (
	"net/url"
	"time"
)

// 扫描
type ScanResponse struct {
	Results []Response `json:"results"`
}

// 响应数据
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

// batch请求数据
type RequestData struct {
	Requests []Requests `json:"requests"`
}

type Requests struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   Body   `json:"body"`
}
type Body struct {
	Content string `json:"content"`
	PubUser string `json:"pubUser"`
}

func (l *LeanCloud) AddClass(name string, data interface{}) (Response, error) {
	r := post(l, "classes/"+name+"?fetchWhenSave=true", data)
	return fetch(r, l)
}

func (l *LeanCloud) GetClass(name, objectId string, data interface{}) (Response, error) {
	r := get(l, "classes/"+name+"/"+objectId, data)
	return fetch(r, l)
}

func (l *LeanCloud) PutClass(name, objectId string, data interface{}) (Response, error) {
	r := put(l, "classes/"+name+"/"+objectId, data)
	return fetch(r, l)
}

func (l *LeanCloud) DeleteClass(name, objectId string, data interface{}) (Response, error) {
	var where string = ""
	switch v := data.(type) {
	case string:
		where = "?where=" + url.QueryEscape(v)
	}
	r := delete(l, "classes/"+name+"/"+objectId+where, nil)
	return fetch(r, l)
}

func (l *LeanCloud) ScanClass(name string, data interface{}) error {
	defer l.mu.Unlock()
	l.mu.Lock()
	l.SetSign(SIGN_MASTER_KEY)
	r := get(l, "scan/classes/"+name, data)
	resp, err := fetchList(r, l)
	if err != nil {
		return err
	}
	l.ScanResponse = resp
	return nil
}

func (l *LeanCloud) BatchUpdate(name, method string, body Body) {
	req := RequestData{
		Requests: make([]Requests, 0, 10),
	}
	switch method {
	case "PUT", "DELETE":
		if len(l.ScanResponse.Results) > 0 {
			for _, i := range l.ScanResponse.Results {
				r := Requests{
					Method: method,
					Path:   "/1.1/classes/" + name + "/" + i.ObjectId,
					Body:   body,
				}
				req.Requests = append(req.Requests, r)
			}
		}
		r := post(l, "batch", req)
		fetch(r, l)
	}
}
