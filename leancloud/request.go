package leancloud

import (
	"fmt"
	"net/http"
)

func (l *LeanCloud) put() {
	r := l.request(http.MethodPut)
}

func (l *LeanCloud) post() {
	r := l.request(http.MethodPost)
}

func (l *LeanCloud) get() {
	r := l.request(http.MethodGet)
}

func (l *LeanCloud) delete() {
	r := l.request(http.MethodDelete)
}


func (l *LeanCloud) request(m string) *http.Request{
	r := &http.Request{
		Method: m,
	}
	l.sign()
	if l.sign_mode == SIGN_APP_KEY {
		r.Header.Add("X-LC-Sign", fmt.Sprintf("%s,%d", l.Sign, l.timestamp))	
	}
	if l.sign_mode == SIGN_MASTER_KEY {
		r.Header.Add("X-LC-Sign", fmt.Sprintf("%s,%d,%s", l.Sign, l.timestamp, "master"))	
	}
	r.Header.Add("X-LC-Id", l.Id)
	return r
}