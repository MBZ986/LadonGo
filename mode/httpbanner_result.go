package mode

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"sync"
)

type HttpBannerResult struct {
	BaseResult
}

func (this *HttpBannerResult) Init(channum ...int) {
	this.ProcessFunc = this.ResultFunc
	if len(channum) > 0 && channum[0] != 0 {
		this.Out = make(chan interface{}, channum[0])
	} else {
		this.Out = make(chan interface{})
	}
	this.Wait = new(sync.WaitGroup)
}

func (this *HttpBannerResult) ResultFunc(o <-chan interface{}) *simplejson.Json {
	j := simplejson.New()
	reslist := make([]map[string]interface{}, 0, 5)
	for v := range o {
		fmt.Println("返回结果", v)
		reslist = append(reslist, v.(map[string]interface{}))
	}
	j.Set("data", reslist)
	return j
}
