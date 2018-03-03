package leancloud

import (
	"fmt"
	"net/url"
	"testing"
)

var objectId string

var lean = NewLeanCould("-gzGzoHsz", "", "")

func Test_NewLeancloud(t *testing.T) {
	lean.SetSign(SIGN_MASTER_KEY)
	r, err := lean.AddClass("test", `{"content": "每个 Java 程序员必备的 8 个开发工具","pubUser": "LeanCloud官方客服","pubTimestamp": 1435541999}`)
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}
	re := r.(Response)
	if int(re.Code) != 0 {
		t.Fail()
	}
	objectId = re.ObjectId
}

func Test_GetClasses(t *testing.T) {
	params := url.Values{}
	lean.SetSign(SIGN_APP_KEY)
	params.Add("include", "author")
	r, err := lean.GetClass("test", objectId, params)
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}
	re := r.(Response)
	if int(re.Code) != 0 {
		t.Fail()
	}
}

func Test_PutClasses(t *testing.T) {
	r, err := lean.PutClass("test", objectId, `{"content": "每个 Golang 程序员必备的 8 个开发工具: http://buzzorange.com/techorange/2015/03/03/9-javascript-ide-editor/"}`)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	re := r.(Response)
	if int(re.Code) != 0 {
		t.Fail()
	}
}

func Test_DeleteClasses(t *testing.T) {
	r, err := lean.DeleteClass("test", objectId, `{"clicks": 0}`)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	re := r.(Response)
	if int(re.Code) != 0 {
		t.Fail()
	}
}
