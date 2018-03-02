package leancloud

import (
	"testing"
)

func Test_NewLeancloud(t *testing.T) {
	b := NewLeanCould("-gzGzoHsz", "", "")
	b.SetSign(SIGN_MASTER_KEY)
	b.AddClass("test", `{"content": "每个 Java 程序员必备的 8 个开发工具","pubUser": "LeanCloud官方客服","pubTimestamp": 1435541999}`)
}
