package leancloud

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	Version = "1.1"
)

var (
	API_URL string
)

const (
	SIGN_APP_KEY = iota
	SIGN_MASTER_KEY
)

type LeanCloud struct {
	mu                    *sync.Mutex
	Id, Key, Sign, Master string
	isMaster              bool
	sign_mode             int
	timestamp             int64
}

func NewLeanCould(id, key, masterKey string) *LeanCloud {
	API_URL = "https://" + id[0:7] + ".api.lncld.net/" + Version + "/"
	return &LeanCloud{
		mu:        &sync.Mutex{},
		Id:        id,
		Key:       key,
		sign_mode: SIGN_APP_KEY,
		Master:    masterKey,
	}
}

// app key SIGN_APP_KEY
// master SIGN_MASTER_KEY
func (l *LeanCloud) SetSign(flag int) {
	l.sign_mode = flag
}

func (l *LeanCloud) sign() {
	var key string
	l.timestamp = time.Now().Unix()
	now := strconv.FormatInt(l.timestamp, 10)
	if l.sign_mode == SIGN_APP_KEY {
		key = l.Key
	}
	if l.sign_mode == SIGN_MASTER_KEY {
		key = l.Master
	}
	b := []byte(now + key)
	l.Sign = fmt.Sprintf("%x", md5.Sum([]byte(b)))
}
