package mode

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"sync"
)

type BaseResult struct {
	Wait        *sync.WaitGroup
	Out         chan interface{}
	ProcessFunc func(o <-chan interface{}) *simplejson.Json
}

func (this *BaseResult) SetWait(wait *sync.WaitGroup) {
	this.Wait = wait
}
func (this *BaseResult) WaitProc() {
	if this.Wait != nil {
		this.Wait.Wait()
	} else {
		panic(fmt.Errorf("同步组为空"))
	}
}
func (this *BaseResult) Done() {
	if this.Wait != nil {
		//fmt.Println("Done")
		this.Wait.Done()
	} else {
		panic(fmt.Errorf("同步组为空"))
	}
}
func (this *BaseResult) CloseOut() {
	close(this.Out)
}
func (this *BaseResult) Add() {
	if this.Wait != nil {
		this.Wait.Add(1)
	} else {
		panic(fmt.Errorf("同步组为空"))
	}
}
func (this *BaseResult) Process() *simplejson.Json {
	if this.ProcessFunc != nil && this.Out != nil {
		return this.ProcessFunc(this.Out)
	}
	panic(fmt.Errorf("处理失败，处理函数或输出队列为初始化"))
}
func (this *BaseResult) Push(data interface{}) {
	if this.Out != nil {
		this.Out <- data
	} else {
		panic(fmt.Errorf("Push失败，管道为初始化"))
	}
}

func (this *BaseResult) SetOutChan(out chan interface{}) {
	this.Out = out
}
func (this *BaseResult) GetOutChan() chan interface{} {
	return this.Out
}

func (this *BaseResult) Init(channum ...int) {
	this.ProcessFunc = this.ResultFunc
	if len(channum) > 0 && channum[0] != 0 {
		this.Out = make(chan interface{}, channum[0])
	} else {
		this.Out = make(chan interface{})
	}
	this.Wait = new(sync.WaitGroup)
}

func (this *BaseResult) ResultFunc(o <-chan interface{}) *simplejson.Json {
	j := simplejson.New()
	reslist := make([]string, 0, 5)
	for v := range o {
		fmt.Println("返回结果", v)
		reslist = append(reslist, v.(string))
	}
	j.Set("data", reslist)
	return j
}
