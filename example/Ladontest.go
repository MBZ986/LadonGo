package main

import (
	"fmt"
	ladon "github.com/MBZ986/LadonGo"
	"github.com/MBZ986/LadonGo/mode"
	"github.com/davecgh/go-spew/spew"
	"sync"
	"time"
)

func main() {
	//time.Sleep(1 * time.Second)
	////go ladon.Run("Ladon","192.168.128.0/24","IcmpScan")
	//ladon.Run("Ladon","192.168.128.0/24","NBTINFO")
	//
	////go ladon.Run("Ladon","FuncList")
	//num := 0
	//fmt.Scanln(&num)
	//byteo := make([]byte, 10)
	//buffer := bytes.NewBuffer(byteo)
	//
	//os.Stdout = buffer
	//if n, err := os.Stdout.Read(bytes);err!=nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(n)
	//}
	//fmt.Printf("捕获到输出：%s\n",string(bytes))
	//testGorouting()

	//testNbtInfo()
	//testPingScan()
	//testIcmpScan()
	//testHttpBanner()
	testHttpTitle()
}

//HttpTitle
func testHttpTitle() {
	proxyRun(new(mode.HttpTitleResult), "Ladon", "192.168.131.0/c", "HttpTitle")
}

//HttpBanner
func testHttpBanner() {
	//var result mode.Result = new(mode.HttpBannerResult)
	//result.Init()
	//go ladon.Run(result, "Ladon", "192.168.131.0/24", "HttpBanner")
	//process := result.Process()
	//result.WaitProc()
	//spew.Dump(process)
	//fmt.Println("545")
	proxyRun(new(mode.HttpBannerResult), "Ladon", "192.168.128.0/c", "HttpBanner")
}

//SnmpScan
func testSnmpScan() {
	proxyRun(new(mode.BaseResult), "Ladon", "192.168.128.0/c", "SnmpScan")
}

//IcmpScan
func testIcmpScan() {
	proxyRun(new(mode.BaseResult), "Ladon", "192.168.128.0/24", "IcmpScan")
}

//NBTINFO
func testNbtInfo() {
	//var result mode.Result = new(mode.NbtInfoResult)
	//result.Init(4)
	//go ladon.Run(result, "Ladon", "192.168.128.0/24", "nbtinfo")
	//process := result.Process()
	//result.WaitProc()
	//spew.Dump(process)

	proxyRun(new(mode.NbtInfoResult), "Ladon", "192.168.128.0/24", "nbtinfo")
}

//PINGSCAN
func testPingScan() {
	//var result mode.Result = new(mode.BaseResult)
	//result.Init()
	//go ladon.Run(result, "Ladon", "192.168.128.0/24", "PingScan")
	//process := result.Process()
	//result.WaitProc()
	//spew.Dump(process)

	proxyRun(new(mode.BaseResult), "Ladon", "192.168.131.0/24", "PingScan")
}

func proxyRun(result mode.Result, params ...string) {
	result.Init()
	go ladon.Run(result, params...)
	process := result.Process()
	//result.WaitProc()
	spew.Dump(process)
}

func resfunc(o <-chan interface{}) {
	for v := range o {
		spew.Dump(v)
	}
}

func testGorouting() {
	bools := make(chan string)
	group := new(sync.WaitGroup)
	group.Add(1)
	go func(x <-chan string) {
		fmt.Println("开始携程")
		for data := range x {
			fmt.Printf("接收到数据\n")
			fmt.Println(data)
		}
		group.Done()
	}(bools)

	for i := 0; i < 10; i++ {
		bools <- fmt.Sprintf("fuck:%d", i)
		time.Sleep(10 * time.Millisecond)
	}
	close(bools)
	group.Wait()
}
