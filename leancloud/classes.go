package leancloud

import (
	"net/url"
)

func (l *LeanCloud) AddClass(name string, data interface{}) (interface{}, error) {
	r := l.post("classes/"+name+"?fetchWhenSave=true", data)
	return fetch(r, l)
}

func (l *LeanCloud) GetClass(name, objectId string, data interface{}) (interface{}, error) {
	r := l.get("classes/"+name+"/"+objectId, data)
	return fetch(r, l)
}

func (l *LeanCloud) PutClass(name, objectId string, data interface{}) (interface{}, error) {
	r := l.put("classes/"+name+"/"+objectId, data)
	return fetch(r, l)
}

func (l *LeanCloud) DeleteClass(name, objectId string, data interface{}) (interface{}, error) {
	var where string = ""
	switch v := data.(type) {
	case string:
		where = "?where=" + url.QueryEscape(v)
	}
	r := l.delete("classes/"+name+"/"+objectId+where, nil)
	return fetch(r, l)
}
