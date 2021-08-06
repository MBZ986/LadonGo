package mode

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

type PingScanResult struct {
	BaseResult
}

func (this *PingScanResult) Push(data interface{}) {
	if this.Out != nil {
		this.Out <- data
	} else {
		panic(fmt.Errorf("Push失败，管道未初始化"))
	}
}

func (this *PingScanResult) ResultFunc(o <-chan interface{}) *simplejson.Json {
	j := simplejson.New()
	reslist := make([]string, 0, 5)
	for v := range o {
		fmt.Println("返回结果", v)
		reslist = append(reslist, v.(string))
	}
	j.Set("data", reslist)
	return j
}
