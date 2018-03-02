package leancloud

func (l *LeanCloud) AddClass(name string, data interface{}) {
	r := l.post("classes/"+name, data)
	fetch(r, l)
}
