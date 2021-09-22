package mode

import (
	"github.com/bitly/go-simplejson"
	"sync"
)

type Result interface {
	Init(channum ...int)
	SetWait(wait *sync.WaitGroup)
	ResultFunc(o <-chan interface{}) *simplejson.Json
	Process() *simplejson.Json
	WaitProc()
	Done()
	CloseOut()
	Add()
	Push(data interface{})
	GetOutChan() chan interface{}
	SetOutChan(out chan interface{})
	SetScanType(scan_type int)
	GetScanType() int
}
