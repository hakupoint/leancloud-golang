package leancloud

import (
	"strconv"
	"time"
	"crypto/md5"
	"io"
	"sync"
)

const (
	Version = "1.1"
	ApiUrl = "https://i5mpgucn.api.lncld.net/" + Version + "/"
)

const (
	SIGN_APP_KEY  = iota
	SIGN_MASTER_KEY
)

type LeanCloud struct {
	mu *sync.Mutex
	Id, Key, Sign, Master string
	isMaster bool
	sign_mode int
	timestamp int64
}

func NewLeanCould (id, key, masterKey string) *LeanCloud{
	return &LeanCloud{
		mu: &sync.Mutex{},
		Id: id,
		Key: key,
		sign_mode: SIGN_APP_KEY,
		Master: masterKey,
	}
}

// app key SIGN_APP_KEY
// master SIGN_MASTER_KEY
func (l *LeanCloud) SetSign(flag int) {
	l.sign_mode = flag
}

func (l *LeanCloud) sign(){
	var key string
	h := md5.New()
	l.timestamp = time.Now().Unix()
	now := strconv.FormatInt(l.timestamp, 10)
	if l.sign_mode == SIGN_APP_KEY {
		key = l.Key
	}
	if l.sign_mode == SIGN_MASTER_KEY {
		key = l.Master
	}
	io.WriteString(h, now + key)
	l.Sign = string(h.Sum(nil))
}